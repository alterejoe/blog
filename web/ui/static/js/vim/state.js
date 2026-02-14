// vim/state.js — shared state

export var SEQUENCE_TIMEOUT = 200;

export var state = {
    mode: "normal", // "insert" | "normal" | "visual"
    target: "input", // "input" | noteIndex (int)
    undoStack: [], // [{ids: [], index}]
    jkBuffer: null,
    nmBuffer: null,
    anchor: null, // visual mode start
    cursor: null, // visual mode end
    inputCursor: 0, // character cursor in input
    noteCursor: 0, // character cursor in note
    desiredCol: 0, // remembered column for j/k movement
    dirtyNotes: {}, // { noteId: true } unsaved edits
};

export var command = {
    count: "",
    operator: "",
    timer: null,
};

export function resetCommand() {
    command.count = "";
    command.operator = "";
    if (command.timer) {
        clearTimeout(command.timer);
        command.timer = null;
    }
}

export function getCount() {
    var c = parseInt(command.count) || 1;
    command.count = "";
    return c;
}
