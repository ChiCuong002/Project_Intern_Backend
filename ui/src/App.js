import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Dashboard from "./pages/Admin/Dashboard";
import Statistics from "./pages/Admin/Statistics";
import Products from "./pages/Admin/Products";
import Users from "./pages/Admin/Users";
import Admin from "./component/Layout/Admin";
import DetailUser from "./pages/Admin/Users/detail.user";
import Login from "./pages/Login/login";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Admin />}>
          <Route index element={<Dashboard />} />
          <Route path="statistics" element={<Statistics />} />
          <Route path="products" element={<Products />} />
          <Route path="users" element={<Users />}/>
          <Route path="users/details/:id" element={<DetailUser />} />
          <Route path="/login" element={<Login/>} />
        </Route>
      </Routes>
    </Router>
  );
}

export default App;
