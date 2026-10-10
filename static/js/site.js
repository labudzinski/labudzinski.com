(function () {
  'use strict';
  var root = document.documentElement;
  var media = window.matchMedia('(prefers-color-scheme: dark)');
  var buttons = document.querySelectorAll('[data-theme-choice]');

  function resolvedTheme() {
    var choice = root.dataset.theme || 'auto';
    return choice === 'auto' ? (media.matches ? 'dark' : 'light') : choice;
  }

  function updateTheme() {
    var choice = root.dataset.theme || 'auto';
    var theme = resolvedTheme();
    buttons.forEach(function (button) {
      button.setAttribute('aria-pressed', String(button.dataset.themeChoice === choice));
    });
    var meta = document.querySelector('meta[name="theme-color"]');
    if (meta) meta.content = theme === 'dark' ? '#14181f' : '#f4f3f0';
    document.querySelectorAll('[data-giscus] script').forEach(function (script) {
      script.dataset.theme = theme;
    });
    var frame = document.querySelector('iframe.giscus-frame');
    if (frame && frame.contentWindow) {
      frame.contentWindow.postMessage({ giscus: { setConfig: { theme: theme } } }, 'https://giscus.app');
    }
  }

  buttons.forEach(function (button) {
    button.addEventListener('click', function () {
      root.dataset.theme = button.dataset.themeChoice;
      try { localStorage.setItem('blog-theme', button.dataset.themeChoice); } catch (_) {}
      updateTheme();
    });
  });
  media.addEventListener('change', updateTheme);
  window.addEventListener('storage', function (event) {
    if (event.key !== 'blog-theme' && event.key !== null) return;
    var choice = event.newValue;
    root.dataset.theme = ['light', 'dark', 'auto'].includes(choice) ? choice : 'auto';
    updateTheme();
  });
  updateTheme();

  var toggle = document.querySelector('.menu-toggle');
  var nav = document.querySelector('#site-navigation');
  function closeMenu() {
    if (!toggle || !nav) return;
    nav.classList.remove('is-open');
    toggle.setAttribute('aria-expanded', 'false');
    toggle.setAttribute('aria-label', 'Otwórz menu');
  }
  if (toggle && nav) {
    toggle.addEventListener('click', function () {
      var open = nav.classList.toggle('is-open');
      toggle.setAttribute('aria-expanded', String(open));
      toggle.setAttribute('aria-label', open ? 'Zamknij menu' : 'Otwórz menu');
    });
    nav.addEventListener('click', function (event) {
      if (event.target.closest('a')) closeMenu();
    });
    document.addEventListener('keydown', function (event) {
      if (event.key === 'Escape' && nav.classList.contains('is-open')) {
        closeMenu();
        toggle.focus();
      }
    });
    window.matchMedia('(min-width: 721px)').addEventListener('change', closeMenu);
  }

  /* Measure the article body rather than the comments or signature below it. */
  var progress = document.querySelector('.reading-progress');
  var readingContent = document.querySelector('[data-reading-content]');
  if (progress && readingContent) {
    var scheduled = false;
    function updateReadingProgress() {
      scheduled = false;
      var rect = readingContent.getBoundingClientRect();
      var headerHeight = document.querySelector('.masthead').getBoundingClientRect().height;
      var viewportHeight = window.innerHeight - headerHeight;
      var length = rect.height - viewportHeight;
      progress.value = length > 0 ? Math.max(0, Math.min(100, (headerHeight - rect.top) / length * 100)) : (rect.bottom <= window.innerHeight ? 100 : 0);
    }
    function scheduleReadingProgress() {
      if (scheduled) return;
      scheduled = true;
      window.requestAnimationFrame(updateReadingProgress);
    }
    window.addEventListener('scroll', scheduleReadingProgress, { passive: true });
    window.addEventListener('resize', scheduleReadingProgress);
    window.addEventListener('load', scheduleReadingProgress);
    if (document.fonts) document.fonts.ready.then(scheduleReadingProgress);
    updateReadingProgress();
  }

  /* Supply the initial theme before Giscus loads, then sync its iframe above. */
  document.querySelectorAll('[data-giscus]').forEach(function (container) {
    var observer = new MutationObserver(function () {
      var frame = container.querySelector('iframe.giscus-frame');
      if (!frame) return;
      frame.addEventListener('load', updateTheme);
      updateTheme();
      observer.disconnect();
    });
    observer.observe(container, { childList: true, subtree: true });
    var script = document.createElement('script');
    Array.from(container.attributes).forEach(function (attribute) {
      if (attribute.name.startsWith('data-') && attribute.name !== 'data-giscus') {
        script.setAttribute(attribute.name, attribute.value);
      }
    });
    script.dataset.theme = resolvedTheme();
    script.src = 'https://giscus.app/client.js';
    script.crossOrigin = 'anonymous';
    script.async = true;
    container.appendChild(script);
  });
})();
