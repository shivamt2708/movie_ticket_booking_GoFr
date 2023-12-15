import { Route, Routes } from "react-router-dom";
import { Login, Signup } from "./pages";
import Home from "./pages/Home";
import AdminHome from "./components/admin/AdminHome"
import SuperAdminHome from "./components/super-admin/SuperAdminHome"
import AddMovie from "./components/super-admin/AddMovie"
import AddMovieHall from "./components/admin/AddMovieHall"
import MyHalls from "./components/admin/MyHalls"
import BookTicket from "./components/admin/BookTicket"

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/admin/book-ticket" element={<BookTicket />} />
        <Route path="/my-halls" element={<MyHalls />} />
        <Route path="/add-movie-hall" element={<AddMovieHall />} />
        <Route path="/add-movie" element={<AddMovie />} />
        <Route path="/admin-home" element={<AdminHome />} />
        <Route path="/super-admin-home" element={<SuperAdminHome />} />
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
      </Routes>
    </div>
  );
}

export default App;