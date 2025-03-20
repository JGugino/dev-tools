import "./style.css";
import App from "./App.svelte";
import { changeView, Views } from "./lib/scripts/view";

changeView(Views.Verify);

const app = new App({
  target: document.getElementById("app"),
});

export default app;
