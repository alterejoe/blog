// static/js/form-behavior/sanitizers.js

FormBehaviors.trim = (input) => {
    input.addEventListener("blur", () => {
        input.value = input.value.trim();
        input.dispatchEvent(new Event("input"));
    });
};

FormBehaviors.removeSpaces = (input) => {
    input.addEventListener("blur", () => {
        input.value = input.value.replace(/\s/g, "");
        input.dispatchEvent(new Event("input"));
    });
};

FormBehaviors.removeSpecialChars = (input) => {
    input.addEventListener("blur", () => {
        input.value = input.value.replace(/[^a-zA-Z0-9_-]/g, "");
        input.dispatchEvent(new Event("input"));
    });
};

FormBehaviors.forceUppercase = (input) => {
    input.addEventListener("blur", () => {
        input.value = input.value.toUpperCase();
        input.dispatchEvent(new Event("input"));
    });
};

FormBehaviors.forceLowercase = (input) => {
    input.addEventListener("blur", () => {
        input.value = input.value.toLowerCase();
        input.dispatchEvent(new Event("input"));
    });
};
