import { Route, Routes } from "react-router-dom";
import "./App.css";
import TicTacToe from "./pages/game";
import { Homepage } from "./pages/homepage";
function App() {
  return (
    <Routes>
      <Route path="/*" element={<Homepage />} />
      <Route path="/game" element={<TicTacToe />} />
    </Routes>
  );
}

export default App;
