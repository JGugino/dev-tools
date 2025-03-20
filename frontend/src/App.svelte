<script lang="ts">
    import { onMount } from "svelte";
    import Footer from "./lib/components/core/Footer.svelte";
    import Header from "./lib/components/core/Header.svelte";
    import VerifyView from "./lib/components/views/VerifyView.svelte";
    import { changeView, getCurrentView, Views } from "./lib/scripts/view";
    import { Toaster } from "svelte-french-toast";

    const title: string = "Dev Tools";
    let currentView: string = getCurrentView();
</script>

<Toaster />

{#if currentView != Views.Verify}
    <Header
        {title}
        on:navTrigger={(e) => {
            changeView(e.detail);
            console.log(getCurrentView());
            currentView = getCurrentView();
        }}
    />
{/if}

<main>
    {#if currentView === Views.Verify}
        <VerifyView
            {title}
            on:verifiedLicense={() => {
                changeView(Views.DashHome);
                currentView = getCurrentView();
            }}
        />
    {:else if currentView === Views.DashHome}
        <h1>Dashboard Home</h1>
    {:else if currentView === Views.DashColor}
        <h1>Color Picker</h1>
    {/if}
</main>
<Footer year={"2025"} version={"0.1"} />
