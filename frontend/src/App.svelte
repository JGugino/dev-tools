<script lang="ts">
    import Footer from "./lib/components/core/Footer.svelte";
    import Header from "./lib/components/core/Header.svelte";
    import VerifyView from "./lib/components/views/VerifyView.svelte";

    const title: string = "Dev Tools";

    const possibleViews: string[] = [
        "verify",
        "dashboard-home",
        "dashboard-color-picker",
    ];

    let currentView: string = possibleViews[0];

    const changeView = (view: string) => {
        if (possibleViews.includes(view)) {
            currentView = view;
        }
    };
</script>

<!-- <Header {title} /> -->
<main>
    <VerifyView {title} />
</main>
<Footer year={"2025"} version={"0.1"} />
