// vim/keymap.js — configurable keybindings

import { SEQUENCE_TIMEOUT } from "./state.js";

// Default keymap: action → key
var defaults = {
    // movement
    moveUp: "k",
    moveDown: "j",
    moveLeft: "h",
    moveRight: "l",
    wordForward: "w",
    wordBack: "b",
    wordForwardBig: "W",
    wordBackBig: "B",
    lineStart: "0",
    lineEnd: "$",
    goLast: "G",
    findChar: "f",
    findCharBack: "F",
    repeatFind: ";",
    repeatFindReverse: ",",
    replaceChar: "r",

    // mode transitions
    insertBefore: "i",
    insertAfter: "a",
    insertLineStart: "I",
    insertLineEnd: "A",
    enterVisual: "v",
    escapeToInput: "Escape",

    // actions
    deleteChar: "x",
    undo: "u",
    save: "s",
    saveAll: "S",
    insertNoteAfter: "o",
    insertNoteBefore: "O",

    // operators (first char of sequence)
    deleteOp: "d",
    goFirstOp: "g",

    // two-char sequences
    exitInsert: "jk",
    exitVisual: "nm",

    // timing (ms)
    sequenceTimeout: 200,
};

// Active keymap (action → key)
export var keymap = {};

// Reverse lookup (key → action) for normal mode single keys
export var keyReverse = {};

// Reverse lookup for visual mode
export var keyReverseVisual = {};

export var ACTIVE_SEQUENCE_TIMEOUT = SEQUENCE_TIMEOUT;

// Build reverse lookups from active keymap
function buildReverse() {
    // Clear and repopulate (keeping same object references)
    var key;
    for (key in keyReverse) delete keyReverse[key];
    for (key in keyReverseVisual) delete keyReverseVisual[key];

    var normalActions = [
        "undo",
        "saveAll",
        "moveUp",
        "moveDown",
        "moveLeft",
        "moveRight",
        "wordForward",
        "wordBack",
        "wordForwardBig",
        "wordBackBig",
        "lineStart",
        "lineEnd",
        "goLast",
        "findChar",
        "findCharBack",
        "repeatFind",
        "repeatFindReverse",
        "insertBefore",
        "insertAfter",
        "insertLineStart",
        "insertLineEnd",
        "enterVisual",
        "escapeToInput",
        "deleteChar",
        "undo",
        "save",
        "insertNoteAfter",
        "insertNoteBefore",
    ];

    var visualActions = [
        "moveUp",
        "moveDown",
        "deleteChar",
        "goLast",
        "enterVisual",
        "escapeToInput",
    ];

    for (var i = 0; i < normalActions.length; i++) {
        var action = normalActions[i];
        var k = keymap[action];
        if (k) keyReverse[k] = action;
    }

    for (var i = 0; i < visualActions.length; i++) {
        var action = visualActions[i];
        var k = keymap[action];
        if (k) keyReverseVisual[k] = action;
    }
}

// Apply defaults then merge overrides
function applyKeymap(overrides) {
    var key;
    // Clear and repopulate (keeping same object reference)
    for (key in keymap) delete keymap[key];
    for (key in defaults) {
        keymap[key] = defaults[key];
    }
    if (overrides) {
        for (key in overrides) {
            if (defaults.hasOwnProperty(key)) {
                keymap[key] = overrides[key];
            }
        }
    }
    ACTIVE_SEQUENCE_TIMEOUT = keymap.sequenceTimeout;
    buildReverse();
}

// Load keymap from server, then call callback
export function loadKeymap(callback) {
    fetch("/pref/vim-keymap")
        .then(function (res) {
            if (res.ok) return res.json();
            return null;
        })
        .then(function (data) {
            applyKeymap(data);
            if (callback) callback();
        })
        .catch(function () {
            applyKeymap(null);
            if (callback) callback();
        });
}

// Save a single binding to server
export function saveKeybind(action, key) {
    keymap[action] = key;
    buildReverse();
    fetch("/pref/vim-keymap/" + action, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ key: key }),
    });
}

// Helper: check if a key matches an action
export function isKey(key, action) {
    return keymap[action] === key;
}

// Helper: get first/second char of a two-char sequence
export function seqFirst(action) {
    var seq = keymap[action];
    return seq ? seq[0] : null;
}

export function seqSecond(action) {
    var seq = keymap[action];
    return seq ? seq[1] : null;
}

// Initialize with defaults immediately (overwritten if server responds)
applyKeymap(null);
