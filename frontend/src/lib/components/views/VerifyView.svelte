<script lang="ts">
    import { onMount, createEventDispatcher } from "svelte";
    import toast from "svelte-french-toast";

    export let title: string;

    const dispatch = createEventDispatcher();

    let licenseKey: string = "";

    const verifyLicense = () => {
        if (licenseKey.length <= 0) {
            toast.error("Please enter a license key");
            return;
        }

        dispatch("verifiedLicense");
    };
</script>

<div id="verify-panel">
    <div class="header">
        <h1>{title}</h1>
    </div>
    <div class="input">
        <label for="license-key-input">License Key</label>
        <input
            type="text"
            id="license-key-input"
            name="license-key-input"
            bind:value={licenseKey}
        />
    </div>
    <button on:click|preventDefault={verifyLicense}>Verify</button>
</div>

<style>
    #verify-panel {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        gap: 1rem;
        margin-top: 6rem;

        & .header {
            font-size: 22pt;
            color: var(--color--gray-three);
            position: relative;
            margin-bottom: 2rem;
            width: 100%;
            text-align: center;

            &::after {
                position: absolute;
                content: "";
                background-color: var(--color--gray-three);
                border-radius: 2px;
                width: 60%;
                height: 6px;
                left: 20%;
                bottom: -1rem;
            }
        }

        & .input {
            display: flex;
            flex-direction: column;
            gap: 0.4rem;

            & label {
                font-size: 14pt;
                color: var(--color--gray-three);
            }

            & input {
                width: 360px;
                height: 50px;
                border: 4px solid var(--color--gray-three);
                border-radius: 2px;
                background-color: var(--color--white-one);
                font-size: 14pt;
                padding: 0 0.8rem;
            }
        }

        & button {
            background-color: var(--color--green-one);
            border: 2px solid var(--color--gray-three);
            width: 160px;
            height: 40px;
            font-size: 14pt;
            border-radius: 2px;
            cursor: pointer;
        }
    }
</style>
