// vim/dom.js — DOM helpers, caching, cursor rendering

import { state, resetCommand } from "./state.js";

// --- Element caches ---

var inputEl = null;
var containerEl = null;
var notesCache = null;

export function invalidateCache() {
    notesCache = null;
}

export function getInput() {
    if (!inputEl) inputEl = document.getElementById("fleeting-post");
    return inputEl;
}

export function getDisplay() {
    return document.getElementById("fleeting-cursor-display");
}

export function getContainer() {
    if (!containerEl)
        containerEl = document.getElementById("fleeting-notes-view");
    return containerEl;
}

export function getNotes() {
    if (!notesCache) {
        notesCache = getContainer().querySelectorAll(".fleeting-note");
    }
    return notesCache;
}

export function noteIdAt(index) {
    var notes = getNotes();
    if (notes[index]) return notes[index].dataset.noteId;
    return null;
}

export function inputAtTop() {
    var input = getInput();
    var container = getContainer();
    return (
        input.compareDocumentPosition(container) &
        Node.DOCUMENT_POSITION_FOLLOWING
    );
}

// --- Input cursor display ---

export function renderInputCursor() {
    var display = getDisplay();
    var input = getInput();
    if (!display) return;

    var val = input.value;
    var pos = state.inputCursor;

    if (pos < 0) pos = 0;
    if (pos > val.length) pos = val.length;
    state.inputCursor = pos;

    var before = val.substring(0, pos);
    var cursorChar = val[pos] || " ";
    var after = val.substring(pos + 1);

    display.innerHTML = "";

    var beforeSpan = document.createElement("span");
    beforeSpan.textContent = before;

    var cursorSpan = document.createElement("span");
    cursorSpan.textContent = cursorChar;
    cursorSpan.className = "bg-white text-black";

    var afterSpan = document.createElement("span");
    afterSpan.textContent = after;

    display.appendChild(beforeSpan);
    display.appendChild(cursorSpan);
    display.appendChild(afterSpan);
}

export function showDisplay() {
    var input = getInput();
    var display = getDisplay();
    if (!display) return;
    input.classList.add("hidden");
    display.classList.remove("hidden");
}

export function hideDisplay() {
    var input = getInput();
    var display = getDisplay();
    if (!display) return;
    display.classList.add("hidden");
    input.classList.remove("hidden");
}

// --- Note cursor display ---

export function renderNoteCursor(index) {
    var notes = getNotes();
    if (!notes[index]) return;

    var p = notes[index].querySelector("p");
    var val = p.textContent;
    var pos = state.noteCursor;

    if (pos < 0) pos = 0;
    if (pos > val.length) pos = val.length;
    state.noteCursor = pos;

    var before = val.substring(0, pos);
    var cursorChar = val[pos] || " ";
    var after = val.substring(pos + 1);

    p.innerHTML = "";

    var beforeSpan = document.createElement("span");
    beforeSpan.textContent = before;

    var cursorSpan = document.createElement("span");
    cursorSpan.textContent = cursorChar;
    cursorSpan.className = "bg-white text-black";

    var afterSpan = document.createElement("span");
    afterSpan.textContent = after;

    p.appendChild(beforeSpan);
    p.appendChild(cursorSpan);
    p.appendChild(afterSpan);
}

export function clearNoteCursor(index) {
    var notes = getNotes();
    if (!notes[index]) return;

    var p = notes[index].querySelector("p");
    p.textContent = p.textContent;
}

// --- Ring ---

export function setRing(index) {
    var notes = getNotes();
    unhighlightInput();
    hideDisplay();

    // clear previous note cursor
    if (typeof state.target === "number" && notes[state.target]) {
        clearNoteCursor(state.target);
    }

    notes.forEach(function (n) {
        n.classList.remove("bg-white/5");
        n.classList.add("bg-quaternary");
    });
    if (notes[index]) {
        notes[index].classList.remove("bg-quaternary");
        notes[index].classList.add("bg-white/5");
        notes[index].scrollIntoView({ block: "nearest" });
    }
    state.target = index;

    // restore desired column, clamped to text length
    var p = notes[index] && notes[index].querySelector("p");
    var len = p ? p.textContent.length : 0;
    state.noteCursor = Math.min(state.desiredCol, Math.max(0, len - 1));

    renderNoteCursor(index);
}

export function clearRing() {
    var notes = getNotes();
    for (var i = 0; i < notes.length; i++) {
        clearNoteCursor(i);
    }
    notes.forEach(function (n) {
        n.classList.remove("bg-white/5");
        n.classList.add("bg-quaternary");
    });
}

// --- Visual highlights ---

export function getVisualRange() {
    if (state.anchor === null || state.cursor === null) return null;
    return {
        start: Math.min(state.anchor, state.cursor),
        end: Math.max(state.anchor, state.cursor),
    };
}

export function renderVisual() {
    var notes = getNotes();
    var range = getVisualRange();
    if (!range) return;

    notes.forEach(function (n, i) {
        n.classList.remove("bg-white/5", "bg-white/10", "border-white");
        n.classList.add("bg-quaternary", "border-transparent");
        if (i >= range.start && i <= range.end) {
            n.classList.remove("bg-quaternary", "border-transparent");
            n.classList.add("border-white", "bg-white/10");
        }
    });

    if (notes[state.cursor]) {
        notes[state.cursor].scrollIntoView({ block: "nearest" });
    }
}

export function clearVisualStyles() {
    getNotes().forEach(function (n) {
        n.classList.remove("border-white", "bg-white/5", "bg-white/10");
        n.classList.add("border-transparent", "bg-quaternary");
    });
}

export function clearVisual() {
    state.anchor = null;
    state.cursor = null;
}

// --- Input highlight ---

export function highlightInput() {
    state.mode = "normal";
    state.target = "input";
    state.jkBuffer = null;
    resetCommand();

    clearRing();
    clearVisualStyles();
    clearVisual();

    var display = getDisplay();
    showDisplay();
    if (display) {
        display.classList.remove("border-primary");
        display.classList.add("border-white");
    }
    renderInputCursor();

    getInput().blur();

    var container = getContainer();
    if (inputAtTop()) {
        container.scrollTop = 0;
    } else {
        container.scrollTop = container.scrollHeight;
    }
}

export function unhighlightInput() {
    var display = getDisplay();
    if (display) {
        display.classList.remove("border-white");
        display.classList.add("border-primary");
    }
    var input = getInput();
    input.classList.remove("border-white");
    input.classList.add("border-primary");
}

// --- Dirty tracking and save ---

export function markDirty() {
    if (state.target === "input") return;
    var notes = getNotes();
    var note = notes[state.target];
    if (!note) return;
    var noteId = note.dataset.noteId;
    state.dirtyNotes[noteId] = true;
    note.classList.add("border-yellow-500");
    note.classList.remove("border-transparent");
}

export function saveAllNotes() {
    var notes = getNotes();
    for (var i = 0; i < notes.length; i++) {
        var noteId = notes[i].dataset.noteId;
        if (state.dirtyNotes[noteId]) {
            saveNote(i);
        }
    }
}

export function saveNote(index) {
    var notes = getNotes();
    if (!notes[index]) return;
    var note = notes[index];
    var noteId = note.dataset.noteId;
    if (!state.dirtyNotes[noteId]) return;

    // clear cursor spans before reading text
    clearNoteCursor(index);

    var p = note.querySelector("p");
    var content = p ? p.textContent : "";

    fetch("/notes/fleeting/" + noteId + "/edit", {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ content: content }),
    })
        .then(function (res) {
            return res.text();
        })
        .then(function (html) {
            var temp = document.createElement("div");
            temp.innerHTML = html.trim();
            var restored = temp.firstElementChild;
            note.replaceWith(restored);
            invalidateCache();
            delete state.dirtyNotes[noteId];

            // re-render cursor on the new element
            if (
                state.mode === "normal" &&
                typeof state.target === "number" &&
                state.target === index
            ) {
                setRing(index);
            }
        });
}
