/* Runs before styles paint; external script keeps the strict CSP intact. */
(function () {
  'use strict';
  var choice = 'auto';
  try { choice = localStorage.getItem('blog-theme') || 'auto'; } catch (_) {}
  if (!['light', 'dark', 'auto'].includes(choice)) choice = 'auto';
  document.documentElement.dataset.theme = choice;
  document.documentElement.classList.add('js');
})();
