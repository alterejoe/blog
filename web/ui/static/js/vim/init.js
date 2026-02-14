// vim/init.js — key dispatcher, HTMX integration, bootstrap

import { state } from "./state.js";
import {
    getInput,
    getContainer,
    getNotes,
    invalidateCache,
    highlightInput,
    setRing,
} from "./dom.js";
import { keymap, loadKeymap } from "./keymap.js";
import {
    handleNormalKey,
    handleVisualKey,
    handleInsertKey,
    hasPendingAction,
} from "./keys.js";
import { enterInsertOnInput, enterInsertOnNote } from "./modes.js";

function onKeyDown(e) {
    var active = document.activeElement;
    var input = getInput();
    var km = keymap;

    console.log(
        "key:",
        e.key,
        "mode:",
        state.mode,
        "target:",
        state.target,
    );

    // nothing relevant focused
    if (
        active !== input &&
        !(active && active.closest && active.closest(".fleeting-note"))
    ) {
        if (state.mode === "visual") {
            handleVisualKey(e);
            return;
        }
        if (e.key === km.insertBefore && !hasPendingAction()) {
            console.log(
                "i intercepted, mode:",
                state.mode,
                "active:",
                document.activeElement,
            );
            e.preventDefault();
            if (
                state.mode === "normal" &&
                typeof state.target === "number"
            ) {
                enterInsertOnNote(state.target);
            } else {
                enterInsertOnInput();
            }
            return;
        }
        if (state.mode === "normal") {
            handleNormalKey(e);
            return;
        }
        return;
    }

    if (state.mode === "insert") {
        handleInsertKey(e);
        return;
    }

    if (state.mode === "visual") {
        handleVisualKey(e);
        return;
    }

    if (state.mode === "normal") {
        if (active && active.tagName === "INPUT" && active !== input)
            return;
        if (active && active.tagName === "TEXTAREA") return;
        handleNormalKey(e);
    }
}

function onAfterSettle() {
    invalidateCache();
    if (state.mode === "normal" && state.target === "input") {
        highlightInput();
        return;
    }
    if (state.mode === "normal" && typeof state.target === "number") {
        var notes = getNotes();
        if (notes.length === 0) {
            highlightInput();
        } else if (state.target >= notes.length) {
            setRing(notes.length - 1);
        } else {
            setRing(state.target);
        }
    }
}

function init() {
    document.addEventListener("keydown", onKeyDown);
    document.body.addEventListener("htmx:afterSettle", onAfterSettle);
    if (getContainer()) {
        highlightInput();
    }
}

// Load keymap from server, then initialize
if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", function () {
        loadKeymap(init);
    });
} else {
    loadKeymap(init);
}
