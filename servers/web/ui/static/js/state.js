document.addEventListener("DOMContentLoaded", () => {
	document.querySelectorAll("[data-dirty-scope]").forEach(initDirtyScope);
});

document.body.addEventListener("htmx:afterSwap", (e) => {
	e.target.querySelectorAll?.("[data-dirty-scope]").forEach(initDirtyScope);
});

function readVal(el) {
	if (el.type === "checkbox" || el.type === "radio") return el.checked;
	return el.value;
}

function initDirtyScope(scope) {
	// Prevent duplicate binding
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
		g.initial.push(readVal(el));

		const handler = () => broadcastGroup(scope, groupId, g);
		el.addEventListener("change", handler);
		el.addEventListener("input", handler);
	});

	// initialize
	for (const [groupId, g] of groups) {
		broadcastGroup(scope, groupId, g);
	}
}

function broadcastGroup(scope, groupId, group) {
	const dirty = group.items.some((el, i) => readVal(el) !== group.initial[i]);
	const evtName = dirty ? `dirty.${groupId}` : `clean.${groupId}`;
	htmx.trigger(scope, evtName, { groupId, dirty });
}
