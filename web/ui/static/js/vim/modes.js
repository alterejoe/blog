// vim/modes.js — mode transitions

import { state, resetCommand } from "./state.js";
import {
    getNotes,
    getInput,
    highlightInput,
    unhighlightInput,
    hideDisplay,
    clearRing,
    clearVisual,
    clearVisualStyles,
    clearNoteCursor,
    setRing,
    invalidateCache,
    renderVisual,
} from "./dom.js";
import { _setEnterNormal as setEnterNormalOnMovement } from "./movement.js";
import { _setEnterInsertOnNote as setEnterInsertOnNoteOnActions } from "./actions.js";

export function enterNormal(index) {
    var notes = getNotes();
    if (notes.length === 0) {
        highlightInput();
        return;
    }

    if (index < 0) index = 0;
    if (index >= notes.length) index = notes.length - 1;

    state.mode = "normal";
    state.nmBuffer = null;
    state.jkBuffer = null;
    clearVisual();
    clearVisualStyles();
    resetCommand();

    getInput().blur();
    hideDisplay();
    setRing(index);
}

export function enterInsertOnInput() {
    state.mode = "insert";
    state.target = "input";
    state.jkBuffer = null;
    clearVisual();
    clearVisualStyles();
    resetCommand();

    unhighlightInput();
    clearRing();
    hideDisplay();

    var input = getInput();
    input.classList.remove("hidden");
    input.focus();
    input.setSelectionRange(state.inputCursor, state.inputCursor);
}

export function enterInsertOnNote(index) {
    var notes = getNotes();
    if (!notes[index]) return;

    var note = notes[index];
    var noteId = note.dataset.noteId;
    var cursorPos = state.noteCursor;

    // clear the rendered cursor before swapping
    clearNoteCursor(index);

    state.mode = "insert";
    state.target = index;
    state.jkBuffer = null;
    clearVisual();
    clearVisualStyles();
    resetCommand();
    hideDisplay();

    fetch("/notes/fleeting/" + noteId + "/edit")
        .then(function (res) {
            return res.text();
        })
        .then(function (html) {
            var temp = document.createElement("div");
            temp.innerHTML = html.trim();
            var editForm = temp.firstElementChild;
            note.replaceWith(editForm);
            invalidateCache();

            var editInput = editForm.querySelector("input");
            if (editInput) {
                editInput.focus();
                editInput.setSelectionRange(cursorPos, cursorPos);
            }
        });
}

function escapeHtml(str) {
    var div = document.createElement("div");
    div.textContent = str;
    return div.innerHTML;
}

export function exitNoteEdit() {
    var notes = getNotes();
    var index = state.target;
    if (typeof index !== "number" || !notes[index]) return;

    var note = notes[index];
    var noteId = note.dataset.noteId;
    var editInput = note.querySelector("input");
    var content = editInput ? editInput.value.trim() : "";

    state.dirtyNotes[noteId] = content;

    var p = note.querySelector("p");
    if (!p && editInput) {
        var timeEl = note.querySelector("[data-utc]");
        var timeHtml = timeEl ? timeEl.outerHTML : "";
        var noteHtml =
            '<div class="fleeting-note p-3 rounded-md text-white text-sm flex flex-row border border-yellow-500/50 bg-quaternary" data-note-id="' +
            noteId +
            '">' +
            "<p>" +
            escapeHtml(content) +
            "</p>" +
            '<span class="flex-1"></span>' +
            timeHtml +
            "</div>";
        var temp = document.createElement("div");
        temp.innerHTML = noteHtml.trim();
        var restored = temp.firstElementChild;
        note.replaceWith(restored);
        invalidateCache();
    }
}

export function enterVisual() {
    if (typeof state.target !== "number") return;

    // clear note cursor before entering visual
    clearNoteCursor(state.target);

    state.mode = "visual";
    state.anchor = state.target;
    state.cursor = state.target;
    resetCommand();
    renderVisual();
}

// --- Wire up circular dependencies ---
setEnterNormalOnMovement(enterNormal);
setEnterInsertOnNoteOnActions(enterInsertOnNote);
