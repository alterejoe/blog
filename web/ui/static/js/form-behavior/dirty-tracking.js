// static/js/form-behavior/dirty-tracking.js

FormBehaviors._readVal = (el) => {
    if (el.type === "checkbox" || el.type === "radio") return el.checked;
    return el.value;
};

FormBehaviors._initDirtyScope = (scope) => {
    if (scope.__dirtyInitialized) return;
    scope.__dirtyInitialized = true;

    const groups = new Map();

    scope.querySelectorAll("[data-dirty-watch]").forEach((el) => {
        const groupId = el.dataset.dirtyGroup || "default";
        if (!groups.has(groupId)) {
            groups.set(groupId, { items: [], initial: [] });
        }
        const g = groups.get(groupId);
        g.items.push(el);
        g.initial.push(FormBehaviors._readVal(el));

        const handler = () => FormBehaviors._broadcastGroup(scope, groupId, g);
        el.addEventListener("change", handler);
        el.addEventListener("input", handler);
    });

    // Initialize all groups
    for (const [groupId, g] of groups) {
        FormBehaviors._broadcastGroup(scope, groupId, g);
    }
};

FormBehaviors._broadcastGroup = (scope, groupId, group) => {
    const dirty = group.items.some(
        (el, i) => FormBehaviors._readVal(el) !== group.initial[i],
    );
    const evtName = dirty ? `dirty.${groupId}` : `clean.${groupId}`;
    htmx.trigger(scope, evtName, { groupId, dirty });
};
