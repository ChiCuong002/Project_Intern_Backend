import { Link, NavLink } from "react-router-dom";
import "./Sidebar.css";

function Sidebar() {
  return (
    <nav className="left-side">
      <ul>
        <li>
          <Link to="/">
            <img src="" alt="" />
          </Link>
        </li>
        <li>
          <NavLink to="/">Dashboard</NavLink>
        </li>
        <li>
          <NavLink to="/statistics">Statistics</NavLink>
        </li>
        <li>
          <NavLink to="/products">Products</NavLink>
        </li>
        <li>
          <NavLink to="/users">Users</NavLink>
        </li>
      </ul>
    </nav>
  );
}
export default Sidebar;
