import { Route, Routes } from "react-router-dom";
import { Login, Signup } from "./pages";
import GetProduct from "./components/GetProduct";

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/get-product" element={<GetProduct />} />
      </Routes>
    </div>
  );
}

export default App;