import { JSX } from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router";

import AuthPage from "./pages/auth/AuthPage";

import "./App.css";

function App() {
  const routes: JSX.Element = (
    <Routes>
      <Route path="/auth" element={<AuthPage />} />
      <Route path="*" element={<Navigate to="/auth" />} />
    </Routes>
  );

  return (
    <>
      <Router>{routes}</Router>
    </>
  );
}

export default App;
