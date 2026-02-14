// vim/actions.js — delete, undo, insert note

import { state } from "./state.js";
import {
    getNotes,
    invalidateCache,
    highlightInput,
    setRing,
    getContainer,
    getVisualRange,
    clearVisual,
    clearVisualStyles,
} from "./dom.js";

// Late-bound ref to enterInsertOnNote (set by modes.js to avoid circular dep)
var enterInsertOnNoteRef = function () {};

export function _setEnterInsertOnNote(fn) {
    enterInsertOnNoteRef = fn;
}

export function deleteNotes(startIndex, count) {
    var notes = getNotes();
    if (typeof startIndex !== "number") return;

    var end = Math.min(startIndex + count - 1, notes.length - 1);
    var ids = [];

    for (var i = end; i >= startIndex; i--) {
        var note = notes[i];
        ids.unshift(note.dataset.noteId);
        note.remove();

        fetch("/notes/fleeting/" + note.dataset.noteId + "/delete", {
            method: "PATCH",
        });
    }

    state.undoStack.push({
        ids: ids,
        index: startIndex,
    });

    invalidateCache();

    var remaining = getNotes();
    if (remaining.length === 0) {
        highlightInput();
        return;
    }

    var newIndex =
        startIndex < remaining.length ? startIndex : remaining.length - 1;
    setRing(newIndex);
}

export function deleteVisualSelection() {
    var range = getVisualRange();
    if (!range) return;

    var count = range.end - range.start + 1;
    state.mode = "normal";
    clearVisual();
    clearVisualStyles();
    deleteNotes(range.start, count);
}

export function undoDelete() {
    if (state.undoStack.length === 0) return;

    var entry = state.undoStack.pop();
    var container = getContainer();

    var promises = entry.ids.map(async function (id) {
        return fetch("/notes/fleeting/" + id + "/restore", {
            method: "PATCH",
        }).then(function (res) {
            return res.text();
        });
    });

    Promise.all(promises).then(function (htmls) {
        var notes = getNotes();
        var insertBefore =
            entry.index < notes.length ? notes[entry.index] : null;

        for (var i = 0; i < htmls.length; i++) {
            var temp = document.createElement("div");
            temp.innerHTML = htmls[i].trim();
            var restored = temp.firstElementChild;

            if (insertBefore) {
                container.insertBefore(restored, insertBefore);
            } else {
                container.appendChild(restored);
            }
        }

        invalidateCache();
        setRing(entry.index);
    });
}

export function insertNote(position) {
    var index = state.target;
    if (typeof index !== "number") return;

    fetch("/notes/fleeting/new", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: "fleeting-post=",
    })
        .then(function (res) {
            return res.text();
        })
        .then(function (html) {
            var temp = document.createElement("div");
            temp.innerHTML = html.trim();
            var newNote = temp.firstElementChild;

            var container = getContainer();
            var notes = getNotes();
            var ref = notes[index];

            if (position === "after") {
                if (ref.nextElementSibling) {
                    container.insertBefore(newNote, ref.nextElementSibling);
                } else {
                    container.appendChild(newNote);
                }
                invalidateCache();
                enterInsertOnNoteRef(index + 1);
            } else {
                container.insertBefore(newNote, ref);
                invalidateCache();
                enterInsertOnNoteRef(index);
            }
        });
}
