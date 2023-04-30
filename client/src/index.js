import React from "react";
import ReactDOM from "react-dom/client";
import App from "./AppClass";
import Hello from "./Hello";

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <App />
    <Hello />
  </React.StrictMode>
);
