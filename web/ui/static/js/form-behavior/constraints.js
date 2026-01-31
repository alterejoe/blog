// static/js/form-behavior/constraints.js

FormBehaviors.required = (input) => {
    const validate = (showError = false) => {
        const isValid = input.value.trim().length > 0;

        FormBehaviors._setValid(
            input,
            isValid,
            isValid ? null : "This field is required",
            showError,
        );
    };

    input.addEventListener("input", () => validate(false));
    input.addEventListener("blur", () => validate(true));
    validate(false);
};

FormBehaviors.email = (input) => {
    const validate = (showError = false) => {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        const value = input.value.trim();
        const isValid = value.length === 0 || emailRegex.test(value);

        FormBehaviors._setValid(
            input,
            isValid,
            isValid ? null : "Must be a valid email address",
            showError,
        );
    };

    input.addEventListener("input", () => validate(false));
    input.addEventListener("blur", () => validate(true));
    validate(false);
};

FormBehaviors.minLength = (input, min) => {
    const validate = (showError = false) => {
        const isValid = input.value.length >= parseInt(min);

        FormBehaviors._setValid(
            input,
            isValid,
            isValid
                ? null
                : `Must be at least ${min} character${min > 1 ? "s" : ""}`,
            showError,
        );
    };

    input.addEventListener("input", () => validate(false));
    input.addEventListener("blur", () => validate(true));
    validate(false);
};

FormBehaviors.maxLength = (input, max) => {
    const validate = (showError = false) => {
        const isValid = input.value.length <= parseInt(max);

        FormBehaviors._setValid(
            input,
            isValid,
            isValid ? null : `Must be no more than ${max} characters`,
            showError,
        );
    };

    input.addEventListener("input", () => validate(false));
    input.addEventListener("blur", () => validate(true));
    validate(false);
};

FormBehaviors.checked = (input) => {
    const validate = (showError = false) => {
        const isValid = input.checked;

        FormBehaviors._setValid(
            input,
            isValid,
            isValid ? null : "This must be checked",
            showError,
        );
    };

    input.addEventListener("change", () => validate(false));
    validate(false);
};

FormBehaviors.alphanumeric = (input) => {
    const validate = (showError = false) => {
        const alphanumericRegex = /^[a-zA-Z0-9_]+$/;
        const value = input.value.trim();
        const isValid = value.length === 0 || alphanumericRegex.test(value);

        FormBehaviors._setValid(
            input,
            isValid,
            isValid
                ? null
                : "Must contain only letters, numbers, and underscores",
            showError,
        );
    };

    input.addEventListener("input", () => validate(false));
    input.addEventListener("blur", () => validate(true));
    validate(false);
};

FormBehaviors.lowercase = (input) => {
    const validate = (showError = false) => {
        const value = input.value;
        const isValid = value === value.toLowerCase();

        FormBehaviors._setValid(
            input,
            isValid,
            isValid ? null : "Must be lowercase",
            showError,
        );
    };

    input.addEventListener("input", () => validate(false));
    input.addEventListener("blur", () => validate(true));
    validate(false);
};

FormBehaviors.noSpaces = (input) => {
    const validate = (showError = false) => {
        const isValid = !input.value.includes(" ");

        FormBehaviors._setValid(
            input,
            isValid,
            isValid ? null : "Must not contain spaces",
            showError,
        );
    };

    input.addEventListener("input", () => validate(false));
    input.addEventListener("blur", () => validate(true));
    validate(false);
};
