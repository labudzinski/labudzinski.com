(function () {
  var meta = document.querySelector('meta[name="ga-id"]');
  if (!meta || !meta.content) {
    return;
  }

  window.dataLayer = window.dataLayer || [];
  function gtag() {
    dataLayer.push(arguments);
  }
  window.gtag = gtag;
  gtag("js", new Date());
  gtag("config", meta.content);
})();
