(function () {
  var el = document.querySelector("[data-lead-uptime]");
  if (!el) {
    return;
  }

  var careerStartYear = parseInt(el.dataset.careerStart, 10);
  if (!careerStartYear) {
    careerStartYear = 2002;
  }

  var careerStart = new Date(careerStartYear + "-01-01T00:00:00Z");

  function plural(count, singular, pluralForm) {
    return count === 1 ? singular : pluralForm;
  }

  function pad2(value) {
    return String(value).padStart(2, "0");
  }

  function formatUptimeRest(now) {
    var time =
      pad2(now.getHours()) +
      ":" +
      pad2(now.getMinutes()) +
      ":" +
      pad2(now.getSeconds());

    var days = Math.floor((now.getTime() - careerStart.getTime()) / 86400000);
    var years = Math.floor(days / 365);
    var remDays = days % 365;

    var uptime =
      "up " + years + " " + plural(years, "year", "years");
    if (remDays > 0) {
      uptime +=
        ", " + remDays + " " + plural(remDays, "day", "days");
    }

    return (
      time +
      " " +
      uptime +
      ",  1 user,  load average: 0.42, 0.38, 0.35"
    );
  }

  function tick() {
    el.textContent = formatUptimeRest(new Date());
  }

  tick();
  setInterval(tick, 1000);
})();
