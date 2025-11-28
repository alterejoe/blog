// htmx.logAll();
function initCheckboxes() {
    const tables = document.querySelectorAll("#table-layout");
    console.log("tables", tables);
    tables.forEach((table) => {
        const checkboxes = table.querySelectorAll('input[type="checkbox"]');
        let lastChecked;
        checkboxes.forEach((box, i) => {
            // Listen on the checkbox itself, not the parent
            box.addEventListener("click", (e) => {
                if (!lastChecked) {
                    lastChecked = box;
                    return;
                }
                if (e.shiftKey) {
                    let start = [...checkboxes].indexOf(lastChecked);
                    let end = i;
                    let [from, to] = [
                        Math.min(start, end),
                        Math.max(start, end),
                    ];
                    for (let j = from; j <= to; j++) {
                        checkboxes[j].checked = box.checked;
                    }
                }
                lastChecked = box;
            });

            // Prevent text selection on shift+mousedown for the label
            box.parentElement.addEventListener("mousedown", (e) => {
                if (e.shiftKey) {
                    e.preventDefault();
                }
            });
        });
    });
}
// Run on initial page load
document.addEventListener("htmx:afterSettle", initCheckboxes);
