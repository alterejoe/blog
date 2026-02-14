// vim/keys.js — key handlers (keymap-driven)

import { state, command, resetCommand, getCount } from "./state.js";
import {
    getNotes,
    getInput,
    highlightInput,
    renderInputCursor,
    renderNoteCursor,
    setRing,
    renderVisual,
    markDirty,
    saveNote,
    saveAllNotes,
    invalidateCache,
} from "./dom.js";
import {
    keymap,
    keyReverse,
    keyReverseVisual,
    ACTIVE_SEQUENCE_TIMEOUT,
    seqFirst,
    seqSecond,
} from "./keymap.js";
import { moveUp, moveDown, goToFirst, goToLast, visualMoveUp, visualMoveDown } from "./movement.js";
import { deleteNotes, deleteVisualSelection, insertNote } from "./actions.js";
import {
    enterNormal,
    enterInsertOnInput,
    enterInsertOnNote,
    exitNoteEdit,
    enterVisual,
} from "./modes.js";

export var _pendingDeleteFind = false;

var pendingF = false;
var lastFChar = null;
var lastFDir = 1;
var pendingFCount = 0;
var pendingReplace = false;
var pendingReplaceCount = 0;

// --- Global undo/redo stacks ---

var undoStack = []; // [{ noteId, text }, ...]
var redoStack = [];

function findNoteIndex(noteId) {
    var notes = getNotes();
    for (var i = 0; i < notes.length; i++) {
        if (notes[i].dataset.noteId === noteId) return i;
    }
    return -1;
}

function pushUndo() {
    if (state.target === "input") return;
    var notes = getNotes();
    var note = notes[state.target];
    if (!note) return;
    var p = note.querySelector("p");
    if (!p) return;
    undoStack.push({
        noteId: note.dataset.noteId,
        text: p.textContent,
    });
    redoStack = [];
}

function updateDirtyState(noteId) {
    var hasDirty = false;
    for (var i = 0; i < undoStack.length; i++) {
        if (undoStack[i].noteId === noteId) {
            hasDirty = true;
            break;
        }
    }
    var index = findNoteIndex(noteId);
    if (index === -1) return;
    var notes = getNotes();
    if (!hasDirty) {
        delete state.dirtyNotes[noteId];
        notes[index].classList.remove("border-yellow-500");
        notes[index].classList.add("border-transparent");
    } else {
        state.dirtyNotes[noteId] = true;
        notes[index].classList.add("border-yellow-500");
        notes[index].classList.remove("border-transparent");
    }
}

export function undoTextEdit() {
    if (undoStack.length === 0) return;
    var entry = undoStack.pop();

    var index = findNoteIndex(entry.noteId);
    if (index === -1) return;

    var notes = getNotes();
    var p = notes[index].querySelector("p");

    // push current state to redo before restoring
    redoStack.push({
        noteId: entry.noteId,
        text: p.textContent,
    });

    p.textContent = entry.text;
    setRing(index);
    updateDirtyState(entry.noteId);
}

export function redoTextEdit() {
    if (redoStack.length === 0) return;
    var entry = redoStack.pop();

    var index = findNoteIndex(entry.noteId);
    if (index === -1) return;

    var notes = getNotes();
    var p = notes[index].querySelector("p");

    // push current state back to undo
    undoStack.push({
        noteId: entry.noteId,
        text: p.textContent,
    });

    p.textContent = entry.text;
    setRing(index);
    markDirty();
}

// --- Helper: get text for current target ---

function getTargetText() {
    if (state.target === "input") {
        return getInput().value;
    }
    var notes = getNotes();
    if (notes[state.target]) {
        return notes[state.target].querySelector("p").textContent;
    }
    return "";
}

function getCursorPos() {
    return state.target === "input"
        ? state.inputCursor
        : state.noteCursor;
}

function setCursorPos(pos) {
    if (state.target === "input") {
        state.inputCursor = pos;
        state.desiredCol = pos;
        renderInputCursor();
    } else {
        state.noteCursor = pos;
        state.desiredCol = pos;
        renderNoteCursor(state.target);
    }
}

// --- Text edit helpers ---

function deleteRange(from, to) {
    var start = Math.min(from, to);
    var stop = Math.max(from, to);
    if (start === stop) return;

    pushUndo();

    if (state.target === "input") {
        var input = getInput();
        var val = input.value;
        input.value = val.substring(0, start) + val.substring(stop);
        state.inputCursor = start;
        if (
            state.inputCursor >= input.value.length &&
            state.inputCursor > 0
        ) {
            state.inputCursor = input.value.length - 1;
        }
        state.desiredCol = state.inputCursor;
        renderInputCursor();
    } else {
        var notes = getNotes();
        var p =
            notes[state.target] &&
            notes[state.target].querySelector("p");
        if (p) {
            var val = p.textContent;
            p.textContent = val.substring(0, start) + val.substring(stop);
            state.noteCursor = start;
            if (
                state.noteCursor >= p.textContent.length &&
                state.noteCursor > 0
            ) {
                state.noteCursor = p.textContent.length - 1;
            }
            state.desiredCol = state.noteCursor;
            renderNoteCursor(state.target);
        }
    }
    markDirty();
}

function deleteAtCursor(count) {
    pushUndo();

    if (state.target === "input") {
        var input = getInput();
        var val = input.value;
        var pos = state.inputCursor;
        for (var xi = 0; xi < count; xi++) {
            if (val.length > 0 && pos < val.length) {
                val = val.substring(0, pos) + val.substring(pos + 1);
            }
        }
        input.value = val;
        if (
            state.inputCursor >= input.value.length &&
            state.inputCursor > 0
        ) {
            state.inputCursor = input.value.length - 1;
        }
        renderInputCursor();
    } else {
        var notes = getNotes();
        var p =
            notes[state.target] &&
            notes[state.target].querySelector("p");
        if (p) {
            var val = p.textContent;
            var pos = state.noteCursor;
            for (var xi = 0; xi < count; xi++) {
                if (val.length > 0 && pos < val.length) {
                    val = val.substring(0, pos) + val.substring(pos + 1);
                }
            }
            p.textContent = val;
            if (
                state.noteCursor >= p.textContent.length &&
                state.noteCursor > 0
            ) {
                state.noteCursor = p.textContent.length - 1;
            }
            renderNoteCursor(state.target);
        }
    }
    markDirty();
}

// --- Word motion helpers ---

function wordForward(val, pos) {
    if (pos < val.length) {
        var isWord = /\w/.test(val[pos]);
        if (isWord) {
            while (pos < val.length && /\w/.test(val[pos])) pos++;
        } else {
            while (
                pos < val.length &&
                !/\w/.test(val[pos]) &&
                !/\s/.test(val[pos])
            )
                pos++;
        }
        while (pos < val.length && /\s/.test(val[pos])) pos++;
    }
    return pos;
}

function wordBack(val, pos) {
    pos = pos - 1;
    while (pos > 0 && /\s/.test(val[pos])) pos--;
    if (pos > 0) {
        var isWord = /\w/.test(val[pos]);
        if (isWord) {
            while (pos > 0 && /\w/.test(val[pos - 1])) pos--;
        } else {
            while (
                pos > 0 &&
                !/\w/.test(val[pos - 1]) &&
                !/\s/.test(val[pos - 1])
            )
                pos--;
        }
    }
    return Math.max(0, pos);
}

function wordForwardBig(val, pos) {
    while (pos < val.length && !/\s/.test(val[pos])) pos++;
    while (pos < val.length && /\s/.test(val[pos])) pos++;
    return pos;
}

function wordBackBig(val, pos) {
    pos = pos - 1;
    while (pos > 0 && /\s/.test(val[pos])) pos--;
    while (pos > 0 && !/\s/.test(val[pos - 1])) pos--;
    return Math.max(0, pos);
}

// --- Find char helpers ---

function findForward(val, pos, ch) {
    return val.indexOf(ch, pos + 1);
}

function findBackward(val, pos, ch) {
    return val.lastIndexOf(ch, pos - 1);
}

// --- Pending action check ---

export function hasPendingAction() {
    return pendingF || pendingReplace || command.operator !== "";
}

// --- Normal mode ---

export function handleNormalKey(e) {
    var key = e.key;
    var cmd = command;
    var km = keymap;

    // Ctrl+R — redo
    if (e.ctrlKey && key === "r") {
        e.preventDefault();
        redoTextEdit();
        return;
    }
    if (pendingReplace) {
        if (
            key === "Shift" ||
            key === "Control" ||
            key === "Alt" ||
            key === "Meta"
        ) {
            return;
        }
        e.preventDefault();
        pendingReplace = false;
        pushUndo();

        if (state.target === "input") {
            var input = getInput();
            var val = input.value;
            var pos = state.inputCursor;
            for (var ri = 0; ri < pendingReplaceCount; ri++) {
                if (pos + ri < val.length) {
                    val =
                        val.substring(0, pos + ri) +
                        key +
                        val.substring(pos + ri + 1);
                }
            }
            input.value = val;
            state.inputCursor = pos + pendingReplaceCount - 1;
            renderInputCursor();
        } else {
            var notes = getNotes();
            var p =
                notes[state.target] &&
                notes[state.target].querySelector("p");
            if (p) {
                var val = p.textContent;
                var pos = state.noteCursor;
                for (var ri = 0; ri < pendingReplaceCount; ri++) {
                    if (pos + ri < val.length) {
                        val =
                            val.substring(0, pos + ri) +
                            key +
                            val.substring(pos + ri + 1);
                    }
                }
                p.textContent = val;
                state.noteCursor = pos + pendingReplaceCount - 1;
                renderNoteCursor(state.target);
            }
        }
        markDirty();
        return;
    }

    // f{char} pending — next key is the target char
    if (pendingF) {
        if (
            key === "Shift" ||
            key === "Control" ||
            key === "Alt" ||
            key === "Meta"
        ) {
            return;
        }
        e.preventDefault();
        pendingF = false;
        lastFChar = key;
        var val = getTargetText();
        var pos = getCursorPos();
        var idx = pos;

        for (var fi = 0; fi < pendingFCount; fi++) {
            if (lastFDir === 1) {
                idx = findForward(val, idx, key);
            } else {
                idx = findBackward(val, idx, key);
            }
            if (idx === -1) break;
        }

        if (idx !== -1) {
            if (_pendingDeleteFind) {
                _pendingDeleteFind = false;
                if (lastFDir === 1) {
                    deleteRange(pos, idx + 1);
                } else {
                    deleteRange(idx, pos);
                }
            } else {
                setCursorPos(idx);
            }
        } else {
            _pendingDeleteFind = false;
        }
        return;
    }

    // ignore bare modifier keys
    if (
        key === "Shift" ||
        key === "Control" ||
        key === "Alt" ||
        key === "Meta"
    ) {
        return;
    }

    // digits → count buffer
    if (key >= "1" && key <= "9") {
        e.preventDefault();
        cmd.count += key;
        return;
    }
    if (key === "0" && cmd.count.length > 0) {
        e.preventDefault();
        cmd.count += key;
        return;
    }

    // f/F — consume count before getCount() clears it
    if (key === km.findChar) {
        e.preventDefault();
        pendingFCount = parseInt(cmd.count) || 1;
        cmd.count = "";
        pendingF = true;
        lastFDir = 1;
        return;
    }

    if (key === km.findCharBack) {
        e.preventDefault();
        pendingFCount = parseInt(cmd.count) || 1;
        cmd.count = "";
        pendingF = true;
        lastFDir = -1;
        return;
    }
    if (key === km.replaceChar) {
        e.preventDefault();
        pendingReplaceCount = parseInt(cmd.count) || 1;
        cmd.count = "";
        pendingReplace = true;
        return;
    }

    // operator: dd (delete line)
    if (key === km.deleteOp) {
        e.preventDefault();
        if (cmd.operator === km.deleteOp) {
            var count = getCount();
            cmd.operator = "";
            if (cmd.timer) {
                clearTimeout(cmd.timer);
                cmd.timer = null;
            }
            if (state.target === "input") {
                getInput().value = "";
                state.inputCursor = 0;
                renderInputCursor();
            } else {
                deleteNotes(state.target, count);
            }
            return;
        }
        cmd.operator = km.deleteOp;
        if (cmd.timer) clearTimeout(cmd.timer);
        cmd.timer = setTimeout(function () {
            cmd.operator = "";
            cmd.timer = null;
        }, ACTIVE_SEQUENCE_TIMEOUT);
        return;
    }

    // operator: gg (go first)
    if (key === km.goFirstOp) {
        e.preventDefault();
        if (cmd.operator === km.goFirstOp) {
            cmd.operator = "";
            if (cmd.timer) {
                clearTimeout(cmd.timer);
                cmd.timer = null;
            }
            goToFirst();
            resetCommand();
            return;
        }
        cmd.operator = km.goFirstOp;
        if (cmd.timer) clearTimeout(cmd.timer);
        cmd.timer = setTimeout(function () {
            cmd.operator = "";
            cmd.timer = null;
        }, ACTIVE_SEQUENCE_TIMEOUT);
        return;
    }

    // d + motion (operator is d, key is the motion)
    if (cmd.operator === km.deleteOp) {
        e.preventDefault();
        cmd.operator = "";
        if (cmd.timer) {
            clearTimeout(cmd.timer);
            cmd.timer = null;
        }

        var val = getTargetText();
        var pos = getCursorPos();
        var count = getCount();
        var end = pos;

        // df{char} / dF{char}
        if (key === km.findChar || key === km.findCharBack) {
            pendingF = true;
            pendingFCount = count;
            lastFDir = key === km.findChar ? 1 : -1;
            _pendingDeleteFind = true;
            return;
        }

        var motionAction = keyReverse[key];

        switch (motionAction) {
            case "wordForward":
                for (var i = 0; i < count; i++) end = wordForward(val, end);
                break;
            case "wordForwardBig":
                for (var i = 0; i < count; i++)
                    end = wordForwardBig(val, end);
                break;
            case "wordBack":
                for (var i = 0; i < count; i++) end = wordBack(val, end);
                break;
            case "wordBackBig":
                for (var i = 0; i < count; i++) end = wordBackBig(val, end);
                break;
            case "lineEnd":
                end = val.length;
                break;
            case "lineStart":
                end = 0;
                break;
            case "moveRight":
                end = Math.min(val.length, pos + count);
                break;
            case "moveLeft":
                end = Math.max(0, pos - count);
                break;
            default:
                resetCommand();
                return;
        }

        deleteRange(pos, end);
        return;
    }

    // all other keys clear operator
    cmd.operator = "";
    if (cmd.timer) {
        clearTimeout(cmd.timer);
        cmd.timer = null;
    }

    var count = getCount();
    var val, pos;

    // use reverse lookup
    var action = keyReverse[key];

    switch (action) {
        case "save":
            e.preventDefault();
            if (typeof state.target === "number") {
                saveNote(state.target);
            }
            break;

        case "saveAll":
            e.preventDefault();
            saveAllNotes();
            break;

        case "moveDown":
            e.preventDefault();
            moveDown(count);
            break;

        case "moveUp":
            e.preventDefault();
            moveUp(count);
            break;

        case "moveLeft":
            e.preventDefault();
            pos = getCursorPos();
            setCursorPos(Math.max(0, pos - count));
            break;

        case "moveRight":
            e.preventDefault();
            val = getTargetText();
            pos = getCursorPos();
            setCursorPos(Math.min(val.length, pos + count));
            break;

        case "wordForward":
            e.preventDefault();
            val = getTargetText();
            pos = getCursorPos();
            for (var wi = 0; wi < count; wi++) {
                pos = wordForward(val, pos);
            }
            setCursorPos(pos);
            break;

        case "wordForwardBig":
            e.preventDefault();
            val = getTargetText();
            pos = getCursorPos();
            for (var wi = 0; wi < count; wi++) {
                pos = wordForwardBig(val, pos);
            }
            setCursorPos(pos);
            break;

        case "wordBack":
            e.preventDefault();
            val = getTargetText();
            pos = getCursorPos();
            for (var bi = 0; bi < count; bi++) {
                pos = wordBack(val, pos);
            }
            setCursorPos(pos);
            break;

        case "wordBackBig":
            e.preventDefault();
            val = getTargetText();
            pos = getCursorPos();
            for (var bi = 0; bi < count; bi++) {
                pos = wordBackBig(val, pos);
            }
            setCursorPos(pos);
            break;

        case "lineStart":
            e.preventDefault();
            setCursorPos(0);
            break;

        case "lineEnd":
            e.preventDefault();
            val = getTargetText();
            setCursorPos(Math.max(0, val.length - 1));
            break;

        case "goLast":
            e.preventDefault();
            goToLast();
            break;

        case "repeatFind":
            e.preventDefault();
            if (lastFChar) {
                val = getTargetText();
                pos = getCursorPos();
                var idx = pos;
                for (var fi = 0; fi < count; fi++) {
                    if (lastFDir === 1) {
                        idx = findForward(val, idx, lastFChar);
                    } else {
                        idx = findBackward(val, idx, lastFChar);
                    }
                    if (idx === -1) break;
                }
                if (idx !== -1) setCursorPos(idx);
            }
            break;

        case "repeatFindReverse":
            e.preventDefault();
            if (lastFChar) {
                val = getTargetText();
                pos = getCursorPos();
                var idx = pos;
                for (var fi = 0; fi < count; fi++) {
                    if (lastFDir === 1) {
                        idx = findBackward(val, idx, lastFChar);
                    } else {
                        idx = findForward(val, idx, lastFChar);
                    }
                    if (idx === -1) break;
                }
                if (idx !== -1) setCursorPos(idx);
            }
            break;

        case "deleteChar":
            e.preventDefault();
            deleteAtCursor(count);
            break;

        case "insertBefore":
            e.preventDefault();
            if (state.target === "input") {
                enterInsertOnInput();
            } else {
                enterInsertOnNote(state.target);
            }
            break;

        case "insertAfter":
            e.preventDefault();
            if (state.target === "input") {
                state.inputCursor = Math.min(
                    getInput().value.length,
                    state.inputCursor + 1,
                );
                enterInsertOnInput();
            } else {
                val = getTargetText();
                state.noteCursor = Math.min(
                    val.length,
                    state.noteCursor + 1,
                );
                enterInsertOnNote(state.target);
            }
            break;

        case "insertLineEnd":
            e.preventDefault();
            if (state.target === "input") {
                state.inputCursor = getInput().value.length;
                enterInsertOnInput();
            } else {
                val = getTargetText();
                state.noteCursor = val.length;
                enterInsertOnNote(state.target);
            }
            break;

        case "insertLineStart":
            e.preventDefault();
            if (state.target === "input") {
                state.inputCursor = 0;
                enterInsertOnInput();
            } else {
                state.noteCursor = 0;
                enterInsertOnNote(state.target);
            }
            break;

        case "insertNoteAfter":
            e.preventDefault();
            if (state.target !== "input") insertNote("after");
            break;

        case "insertNoteBefore":
            e.preventDefault();
            if (state.target !== "input") insertNote("before");
            break;

        case "undo":
            e.preventDefault();
            undoTextEdit();
            break;

        case "enterVisual":
            e.preventDefault();
            enterVisual();
            break;

        case "escapeToInput":
            e.preventDefault();
            pendingF = false;
            enterInsertOnInput();
            break;
    }
}

// --- Visual mode ---

export function handleVisualKey(e) {
    var key = e.key;
    var cmd = command;
    var km = keymap;

    // exit visual sequence
    var exitFirst = seqFirst("exitVisual");
    var exitSecond = seqSecond("exitVisual");

    if (key === exitFirst) {
        state.nmBuffer = Date.now();
        e.preventDefault();
        return;
    }

    if (
        key === exitSecond &&
        state.nmBuffer &&
        Date.now() - state.nmBuffer < ACTIVE_SEQUENCE_TIMEOUT
    ) {
        state.nmBuffer = null;
        e.preventDefault();
        enterNormal(state.cursor);
        return;
    }
    state.nmBuffer = null;

    // ignore bare modifier keys
    if (
        key === "Shift" ||
        key === "Control" ||
        key === "Alt" ||
        key === "Meta"
    ) {
        return;
    }

    // digits → count
    if (key >= "1" && key <= "9") {
        e.preventDefault();
        cmd.count += key;
        return;
    }
    if (key === "0" && cmd.count.length > 0) {
        e.preventDefault();
        cmd.count += key;
        return;
    }

    var count = getCount();

    // goFirst operator in visual
    if (key === km.goFirstOp) {
        e.preventDefault();
        if (cmd.operator === km.goFirstOp) {
            cmd.operator = "";
            if (cmd.timer) {
                clearTimeout(cmd.timer);
                cmd.timer = null;
            }
            state.cursor = 0;
            renderVisual();
            return;
        }
        cmd.operator = km.goFirstOp;
        if (cmd.timer) clearTimeout(cmd.timer);
        cmd.timer = setTimeout(function () {
            cmd.operator = "";
            cmd.timer = null;
        }, ACTIVE_SEQUENCE_TIMEOUT);
        return;
    }

    if (key === km.moveDown) {
        e.preventDefault();
        visualMoveDown(count);
    } else if (key === km.moveUp) {
        e.preventDefault();
        visualMoveUp(count);
    } else if (key === km.deleteOp || key === km.deleteChar) {
        e.preventDefault();
        deleteVisualSelection();
    } else if (key === km.goLast) {
        e.preventDefault();
        state.cursor = getNotes().length - 1;
        renderVisual();
    } else if (key === km.enterVisual || key === "Escape") {
        e.preventDefault();
        enterNormal(state.cursor);
    }
}

// --- Insert mode ---

export function handleInsertKey(e) {
    var key = e.key;
    var km = keymap;

    var exitFirst = seqFirst("exitInsert");
    var exitSecond = seqSecond("exitInsert");

    if (key === "Escape") {
        e.preventDefault();
        if (state.target !== "input") {
            exitNoteEdit();
            enterNormal(state.target);
        } else {
            state.inputCursor = getInput().selectionStart || 0;
            highlightInput();
        }
        return;
    }

    if (key === exitFirst) {
        state.jkBuffer = Date.now();
        return;
    }

    if (
        key === exitSecond &&
        state.jkBuffer &&
        Date.now() - state.jkBuffer < ACTIVE_SEQUENCE_TIMEOUT
    ) {
        state.jkBuffer = null;
        e.preventDefault();

        if (state.target === "input") {
            var input = getInput();
            input.value = input.value.slice(0, -1);
            state.inputCursor = Math.min(
                input.selectionStart || 0,
                input.value.length,
            );
            highlightInput();
        } else {
            var notes = getNotes();
            var editInput =
                notes[state.target] &&
                notes[state.target].querySelector("input");
            if (editInput) {
                editInput.value = editInput.value.slice(0, -1);
            }
            exitNoteEdit();
            enterNormal(state.target);
        }
        return;
    }
}
