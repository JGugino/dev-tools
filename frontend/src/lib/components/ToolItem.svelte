<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import picker from "../../assets/images/color-picker.svg";
    import { Views } from "../scripts/view";

    const dispatch = createEventDispatcher();

    export let toolIcon: string;
    export let toolName: string;
    export let toolView: Views;
</script>

<button
    id="tool-item"
    on:click|preventDefault|stopPropagation={() => {
        dispatch("toolClicked", toolView);
    }}
>
    <div class="icon">
        <img src={toolIcon} alt="color picker icon" />
    </div>

    <h3>{toolName}</h3>
</button>

<style>
    #tool-item {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 1rem;
        width: 96%;
        min-height: 40px;
        padding: 1rem 0;
        background-color: var(--color--gray-one);
        border: 1px solid var(--color--gray-one);
        border-radius: 4px;
        cursor: pointer;

        & h3 {
            font-weight: normal;
            font-size: 15pt;
            color: var(--color--gray-three);
            font-family: var(--font-family--default);
        }

        & .icon {
            display: flex;
            justify-content: center;
            align-items: center;
            --size: 42px;
            width: var(--size);
            height: var(--size);
            background-color: var(--color--gray-three);
            padding: 0.4rem;
            border-radius: 100vh;

            & img {
                --size: 22px;
                width: var(--size);
                height: var(--size);
            }
        }
    }
</style>
