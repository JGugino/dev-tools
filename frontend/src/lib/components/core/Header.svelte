<script lang="ts">
    import { Views } from "../../../lib/scripts/view";
    import pickerIcon from "../../../assets/images/color-picker.svg";

    import ToolItem from "../ToolItem.svelte";
    import { createEventDispatcher } from "svelte";

    export let title: string;

    const dispatch = createEventDispatcher();

    let menuOpen: boolean = true;
</script>

<header>
    <button
        class="header-button"
        on:click={() => {
            dispatch("navTrigger", Views.DashHome);
        }}>{title}</button
    >
    <button
        aria-label="menu button"
        on:click|preventDefault={() => {
            menuOpen = !menuOpen;
        }}
        ><svg
            width="30"
            height="30"
            viewBox="0 0 30 30"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
        >
            <mask
                id="mask0_6_7"
                style="mask-type:alpha"
                maskUnits="userSpaceOnUse"
                x="0"
                y="0"
                width="30"
                height="30"
            >
                <rect width="30" height="30" fill="#D9D9D9" />
            </mask>
            <g mask="url(#mask0_6_7)">
                <path
                    d="M3.75 22.5V20H26.25V22.5H3.75ZM3.75 16.25V13.75H26.25V16.25H3.75ZM3.75 10V7.5H26.25V10H3.75Z"
                    fill="#7E8D85"
                />
            </g>
        </svg>
    </button>
</header>

{#if menuOpen}
    <nav>
        <ToolItem
            toolIcon={pickerIcon}
            toolName={"Color Picker"}
            toolView={Views.DashColor}
            on:toolClicked={(e) => {
                dispatch("navTrigger", e.detail);
                if (menuOpen) {
                    menuOpen = false;
                }
            }}
        />
    </nav>
{/if}

<style>
    header {
        display: flex;
        justify-content: space-between;
        padding: 1.4rem 1rem;
        align-items: center;
        position: relative;
        width: 100%;

        &::after {
            position: absolute;
            content: "";
            background-color: var(--color--gray-three);
            width: 90%;
            height: 6px;
            border-radius: 2px;
            bottom: 0;
            left: 5%;
        }

        & .header-button {
            font-size: 28pt;
            color: var(--color--gray-three);
            user-select: none;
            font-family: var(--font-family--default);
        }

        & button {
            background-color: var(--color--white-one);
            border: 1px solid var(--color--white-one);
            cursor: pointer;

            & svg {
                --size: 38px;
                width: var(--size);
                height: var(--size);
            }
        }
    }

    nav {
        position: fixed;
        right: 0;
        width: 300px;
        height: 100%;
        padding: 2rem;
        &::after {
            position: absolute;
            content: "";
            background-color: var(--color--green-one);
            height: 80%;
            width: 6px;
            left: 0;
            bottom: 15%;
            border-radius: 2px;
        }
    }
</style>
