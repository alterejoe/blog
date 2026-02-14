// vim/movement.js — movement functions

import { state } from "./state.js";
import {
    getNotes,
    inputAtTop,
    highlightInput,
    unhighlightInput,
    setRing,
    renderVisual,
} from "./dom.js";

export function moveUp(count) {
    if (state.target === "input") {
        if (!inputAtTop()) {
            unhighlightInput();
            enterNormalRef(getNotes().length - 1);
        }
        return;
    }
    var target = state.target - count;
    if (target < 0) {
        if (inputAtTop()) {
            highlightInput();
            return;
        }
        target = 0;
    }
    setRing(target);
}

export function moveDown(count) {
    if (state.target === "input") {
        if (inputAtTop()) {
            unhighlightInput();
            enterNormalRef(0);
        }
        return;
    }
    var max = getNotes().length - 1;
    var target = state.target + count;
    if (target > max) {
        if (!inputAtTop()) {
            highlightInput();
            return;
        }
        target = max;
    }
    setRing(target);
}

export function goToFirst() {
    if (inputAtTop()) {
        highlightInput();
    } else {
        unhighlightInput();
        var notes = getNotes();
        if (notes.length > 0) setRing(0);
    }
}

export function goToLast() {
    if (!inputAtTop()) {
        highlightInput();
    } else {
        unhighlightInput();
        var notes = getNotes();
        if (notes.length > 0) setRing(notes.length - 1);
    }
}

export function visualMoveUp(count) {
    var target = state.cursor - count;
    if (target < 0) target = 0;
    state.cursor = target;
    renderVisual();
}

export function visualMoveDown(count) {
    var max = getNotes().length - 1;
    var target = state.cursor + count;
    if (target > max) target = max;
    state.cursor = target;
    renderVisual();
}

// --- Circular dependency resolution ---
// movement needs enterNormal from modes, but modes needs movement indirectly.
// We use a late-bound reference that modes.js sets.
var enterNormalRef = function () {};

export function _setEnterNormal(fn) {
    enterNormalRef = fn;
}
