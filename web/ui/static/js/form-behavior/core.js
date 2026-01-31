// static/js/form-behavior/core.js

const FormBehaviors = {
    _validInputs: new Set(),

    _setValid: (input, isValid, errorMessage = null, showError = true) => {
        if (isValid) {
            FormBehaviors._validInputs.add(input);
            input.classList.remove("invalid");
            input.classList.add("valid");
            FormBehaviors._clearError(input);
        } else {
            FormBehaviors._validInputs.delete(input);
            input.classList.remove("valid");
            input.classList.add("invalid");
            if (errorMessage && showError) {
                FormBehaviors._showError(input, errorMessage);
            }
        }
        FormBehaviors._checkFormValidity(input);
    },

    _showError: (input, message) => {
        const errorDiv = document.getElementById(input.id + "-errors");
        if (errorDiv) {
            errorDiv.textContent = message;
            errorDiv.classList.remove("hidden");
        }
    },

    _clearError: (input) => {
        const errorDiv = document.getElementById(input.id + "-errors");
        if (errorDiv) {
            errorDiv.textContent = "";
            errorDiv.classList.add("hidden");
        }
    },

    _checkFormValidity: (input) => {
        const form = input.closest("[data-constraint-form]");
        if (!form) return;

        const requiredInputs = form.querySelectorAll(
            "[data-constraint][required]",
        );

        const allValid = Array.from(requiredInputs).every((input) => {
            if (!input.value.trim() && input.type !== "checkbox") return false;
            return FormBehaviors._validInputs.has(input);
        });

        const formId = form.id;

        // Look for checkbox inside form OR outside with form attribute
        let checkbox = form.querySelector("[data-enable-on-valid]");
        if (!checkbox && formId) {
            checkbox = document.querySelector(
                `[data-enable-on-valid][form="${formId}"]`,
            );
        }

        if (checkbox) {
            // Checkbox gate pattern
            checkbox.disabled = !allValid;

            if (!checkbox.disabled && checkbox.checked) {
                FormBehaviors._updateEnableTarget(checkbox, formId);
            }
        } else {
            // Direct submit button pattern
            let submitBtn = form.querySelector('button[type="submit"]');
            if (!submitBtn && formId) {
                submitBtn = document.querySelector(
                    `button[type="submit"][form="${formId}"]`,
                );
            }

            if (submitBtn) {
                submitBtn.disabled = !allValid;
            }
        }
    },

    _updateEnableTarget: (checkbox, formId) => {
        const targetSelector = checkbox.dataset.enableTarget;
        if (!targetSelector) return;

        const form = formId
            ? document.getElementById(formId)
            : checkbox.closest("form");
        let target = form ? form.querySelector(targetSelector) : null;

        if (!target) {
            target = document.querySelector(targetSelector);
        }

        if (target) {
            target.disabled = !checkbox.checked;
        }
    },
};
