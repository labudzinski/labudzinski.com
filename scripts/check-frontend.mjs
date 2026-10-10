/* Offline privacy regression checks: no test traffic is sent to Google. */
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import vm from 'node:vm';
import { test } from 'node:test';

const source = readFileSync(new URL('../static/js/gtag-init.js', import.meta.url), 'utf8');
const key = 'blog-analytics-consent-v1';
const id = 'G-TESTLOCAL';
const age = 180 * 24 * 60 * 60 * 1000;
const record = (analytics, savedAt = Date.now()) => JSON.stringify({version: 1, analytics, savedAt});

function harness({ saved, brokenStorage = false, brokenWrites = false } = {}) {
  const events = new Map();
  const scripts = [];
  const writes = [];
  const storage = new Map(saved === undefined ? [] : [[key, saved]]);
  const cookies = new Map([['_ga', 'old'], ['_ga_TESTLOCAL', 'old'], ['site_setting', 'keep']]);
  let reloads = 0;
  const document = { hidden: false, activeElement: null };
  function element(dataset = {}) {
    const listeners = new Map();
    return { dataset, hidden: false, textContent: '',
      addEventListener: (type, cb) => listeners.set(type, cb),
      fire: (type, event = {}) => listeners.get(type)?.(event),
      focus() { document.activeElement = this; }
    };
  }
  const accept = element({cookieChoice: 'granted'});
  const reject = element({cookieChoice: 'denied'});
  const close = element();
  const settings = element();
  const status = element();
  const banner = element();
  banner.hidden = true;
  banner.contains = target => [accept, reject, close].includes(target);
  banner.querySelector = selector => ({
    '[data-cookie-close]': close,
    '[data-cookie-status]': status,
    '[data-cookie-choice="denied"]': reject
  })[selector];
  banner.querySelectorAll = () => [accept, reject];
  document.querySelector = selector => ({
    'meta[name="ga-id"]': {content: id}, '.cookie-banner': banner,
    '[data-cookie-settings]': settings
  })[selector];
  document.querySelectorAll = () => [settings];
  document.createElement = () => ({});
  document.head = {appendChild: script => scripts.push(script)};
  Object.defineProperty(document, 'cookie', {
    get: () => [...cookies].map(([k, v]) => `${k}=${v}`).join('; '),
    set: value => { writes.push(value); cookies.delete(value.split('=')[0]); }
  });
  const addEventListener = (type, cb) => events.set(type, cb);
  document.addEventListener = addEventListener;
  const window = { addEventListener };
  const location = { hostname: 'blog.example.com', pathname: '/posts/example/', reload: () => reloads++ };
  const localStorage = {
    getItem: k => { if (brokenStorage) throw Error('unavailable'); return storage.get(k) ?? null; },
    setItem: (k,v) => { if (brokenStorage || brokenWrites) throw Error('unavailable'); storage.set(k,v); }
  };
  vm.runInNewContext(source, {window, document, localStorage, location, Date, Number});
  return {scripts, window, banner, cookies, writes, storage, accept, reject, settings, close, status,
    fire: (type, event) => events.get(type)?.(event), reloads: () => reloads};
}

test('first visit and rejection make no Google request', () => {
  const h = harness();
  assert.equal(h.banner.hidden, false);
  assert.equal(h.scripts.length, 0);
  assert.equal(h.window['ga-disable-' + id], true);
  assert.equal(h.cookies.has('_ga'), false);
  assert.equal(h.cookies.get('site_setting'), 'keep');
  h.reject.fire('click');
  assert.equal(h.banner.hidden, true);
  assert.equal(JSON.parse(h.storage.get(key)).analytics, 'denied');
  assert.equal(h.scripts.length, 0);
});

test('saved rejection does not reopen the banner or load Google', () => {
  const h = harness({saved: record('denied')});
  assert.equal(h.banner.hidden, true);
  assert.equal(h.scripts.length, 0);
  h.settings.fire('click');
  assert.equal(h.banner.hidden, false);
  assert.match(h.status.textContent, /wyłączona/);
  h.close.fire('click');
  assert.equal(h.banner.hidden, true);
});

test('acceptance loads once, after consent, with advertising denied', () => {
  const h = harness();
  h.accept.fire('click');
  h.accept.fire('click');
  assert.equal(h.scripts.length, 1);
  assert.equal(h.window['ga-disable-' + id], false);
  const calls = h.window.dataLayer.map(args => Array.from(args));
  assert.equal(calls[0][1], 'default');
  assert.equal(calls[0][2].analytics_storage, 'denied');
  assert.equal(calls[1][1], 'update');
  assert.equal(calls[1][2].analytics_storage, 'granted');
  assert.equal(calls[1][2].ad_storage, 'denied');
  assert.equal(calls[1][2].ad_user_data, 'denied');
  assert.equal(calls[1][2].ad_personalization, 'denied');
  const config = calls.find(args => args[0] === 'config')[2];
  assert.equal(config.allow_google_signals, false);
  assert.equal(config.allow_ad_personalization_signals, false);
  assert.equal(config.cookie_update, false);
});

test('saved consent loads without reopening the banner', () => {
  const h = harness({saved: record('granted')});
  assert.equal(h.scripts.length, 1);
  assert.equal(h.banner.hidden, true);
});

test('withdrawal disables the tag, deletes cookies and reloads without the tag', () => {
  const h = harness({saved: record('granted')});
  h.reject.fire('click');
  assert.equal(h.window['ga-disable-' + id], true);
  assert.equal(h.cookies.has('_ga'), false);
  assert.equal(h.cookies.has('_ga_TESTLOCAL'), false);
  assert.ok(h.writes.some(value => value.includes('domain=example.com')));
  assert.equal(h.reloads(), 1);
  const next = harness({saved: h.storage.get(key)});
  assert.equal(next.scripts.length, 0);
});

for (const [label, saved] of [
  ['expired', record('granted', Date.now() - age - 1000)],
  ['future', record('granted', Date.now() + 60000)],
  ['malformed', '{broken'],
  ['wrong version', JSON.stringify({version: 2, analytics: 'granted', savedAt: Date.now()})],
  ['invalid choice', record('yes')]
]) {
  test(`${label} consent stays disabled and asks again`, () => {
    const h = harness({saved});
    assert.equal(h.scripts.length, 0);
    assert.equal(h.banner.hidden, false);
  });
}

test('withdrawal in another tab disables this tab too', () => {
  const h = harness({saved: record('granted')});
  h.storage.set(key, record('denied'));
  h.fire('storage', {key});
  assert.equal(h.window['ga-disable-' + id], true);
  assert.equal(h.reloads(), 1);
});

test('expired consent in a long-open tab stops measurements', () => {
  const h = harness({saved: record('granted')});
  h.storage.set(key, record('granted', Date.now() - age - 1000));
  h.fire('focus');
  assert.equal(h.window['ga-disable-' + id], true);
  assert.equal(h.reloads(), 1);
});

for (const failure of ['brokenStorage', 'brokenWrites']) {
  test(`${failure}: the current choice works in memory without a reload loop`, () => {
    const h = harness({[failure]: true});
    h.accept.fire('click');
    h.fire('focus');
    assert.equal(h.window['ga-disable-' + id], false);
    h.reject.fire('click');
    h.fire('focus');
    assert.equal(h.window['ga-disable-' + id], true);
    assert.equal(h.banner.hidden, true);
    assert.equal(h.reloads(), 0);
  });
}

test('the theme is restored before rendering, including storage errors', () => {
  const source = readFileSync(new URL('../static/js/theme-init.js', import.meta.url), 'utf8');
  for (const [saved, expected] of [['light', 'light'], ['dark', 'dark'], ['auto', 'auto'], ['invalid', 'auto'], [null, 'auto']]) {
    const root = {dataset: {}, classList: {add() {}}};
    vm.runInNewContext(source, {document: {documentElement: root}, localStorage: {getItem: () => saved}});
    assert.equal(root.dataset.theme, expected);
  }
  const root = {dataset: {}, classList: {add() {}}};
  vm.runInNewContext(source, {document: {documentElement: root}, localStorage: {getItem() {throw Error('blocked');}}});
  assert.equal(root.dataset.theme, 'auto');
});
