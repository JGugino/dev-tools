<script lang="ts">
    export let year: string = "2025";
    export let version: string = "0.1";
</script>

<footer>
    <span>&copy; copyright Gugino {year}, All rights reserved.</span>
    <span>dev-tool-v{version}</span>
</footer>

<style>
    footer {
        display: flex;
        justify-content: space-between;
        align-items: center;
        position: fixed;
        bottom: 0;
        width: 100%;
        padding: 0.6rem;
    }
</style>
