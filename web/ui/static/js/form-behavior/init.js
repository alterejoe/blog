// static/js/form-behavior/init.js

const initializeFormBehaviors = (container = document) => {
    // Initialize constraints
    container.querySelectorAll("[data-constraint]").forEach((input) => {
        const constraints = input.dataset.constraint.split(",");
        constraints.forEach((constraint) => {
            const parts = constraint.trim().split(":");
            const name = parts[0];
            const params = parts[1] ? parts[1].split("|") : [];

            if (FormBehaviors[name]) {
                FormBehaviors[name](input, ...params);
            }
        });
    });

    // Initialize dirty tracking
    container.querySelectorAll("[data-dirty-scope]").forEach((scope) => {
        FormBehaviors._initDirtyScope(scope);
    });
};

document.addEventListener("DOMContentLoaded", () => {
    initializeFormBehaviors();
});

document.body.addEventListener("htmx:afterSwap", (e) => {
    initializeFormBehaviors(e.detail.elt);
});

// Auto-enable/disable targets when checkbox state changes
document.addEventListener("change", (e) => {
    if (e.target.hasAttribute("data-enable-on-valid")) {
        FormBehaviors._updateEnableTarget(e.target);
    }
});
