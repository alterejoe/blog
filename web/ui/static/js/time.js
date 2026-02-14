function formatTime(date, format) {
    if (format === "Absolute") {
        return date.toLocaleTimeString([], {
            hour: "numeric",
            minute: "2-digit",
        });
    }
    const now = new Date();
    const diff = Math.floor((now - date) / 1000);
    const days = Math.floor(diff / 86400);
    const hrs = Math.floor((diff % 86400) / 3600);
    const mins = Math.floor((diff % 3600) / 60);
    const secs = diff % 60;
    if (format === "Compact") {
        if (days > 0) return `${days}d ${pad(hrs)}:${pad(mins)}`;
        if (hrs > 0) return `${pad(hrs)}:${pad(mins)}:${pad(secs)}`;
        if (mins > 0) return `${pad(mins)}:${pad(secs)}`;
        return pad(secs);
    }
    if (format === "Verbose") {
        if (days > 0)
            return `${days} day${days > 1 ? "s" : ""} ${hrs} hr${hrs !== 1 ? "s" : ""}`;
        if (hrs > 0)
            return `${hrs} hr${hrs !== 1 ? "s" : ""} ${mins} min${mins !== 1 ? "s" : ""}`;
        if (mins > 0)
            return `${mins} min${mins !== 1 ? "s" : ""} ${secs} sec${secs !== 1 ? "s" : ""}`;
        return `${secs} sec${secs !== 1 ? "s" : ""}`;
    }
}

function pad(n) {
    return n.toString().padStart(2, "0");
}

let timeUpdateInterval = null;

function localizeTimestamps() {
    const format = localStorage.getItem("fleeting_time") || "Absolute";
    document.querySelectorAll("[data-utc]").forEach((el) => {
        const date = new Date(el.dataset.utc);
        el.textContent = formatTime(date, format);
    });

    // Clear existing interval
    if (timeUpdateInterval) {
        clearInterval(timeUpdateInterval);
        timeUpdateInterval = null;
    }

    // Only run interval for relative formats
    if (format === "Compact" || format === "Verbose") {
        timeUpdateInterval = setInterval(() => {
            document.querySelectorAll("[data-utc]").forEach((el) => {
                const date = new Date(el.dataset.utc);
                el.textContent = formatTime(date, format);
            });
        }, 1000);
    }
}

document.addEventListener("htmx:afterRequest", (e) => {
    const match = e.detail.pathInfo?.requestPath?.match(
        /\/pref\/fleeting_time\/(\w+)/,
    );
    if (match) {
        localStorage.setItem("fleeting_time", match[1]);
        localizeTimestamps();
    }
});

document.addEventListener("DOMContentLoaded", localizeTimestamps);
document.addEventListener("htmx:afterSwap", localizeTimestamps);
document.addEventListener("htmx:afterSettle", localizeTimestamps);
