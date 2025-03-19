<script lang="ts">
    export let title: string;
</script>

<header>
    <h1>{title}</h1>
    <button aria-label="menu button"
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

        & h1 {
            font-size: 28pt;
            color: var(--color--gray-three);
            user-select: none;
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
</style>
