import { Link, NavLink } from "react-router-dom";
import "bootstrap-icons/font/bootstrap-icons.css";
// import "bootstrap/dist/css/bootstrap.min.css";
import "./Sidebar.css";
import imageLogo from "~/assets/images/logo_huflit.png";
function Sidebar() {
  return (
    <nav className="left-side">
      <ul>
        <li>
          <Link to="/">
            <img src={imageLogo} alt="logo" />
          </Link>
        </li>
        <li className="left">
          <i className="fa-3x bi bi-house icon"></i>
          <NavLink to="/" className="nav-link">
            Dashboard
          </NavLink>
        </li>
        <li className="left">
          <i className="bi bi-bar-chart icon"></i>
          <NavLink to="/statistics" className="nav-link">
            Statistics
          </NavLink>
        </li>
        <li className="left">
          <i className="bi bi-box icon"></i>
          <NavLink to="/products" className="nav-link">
            Products
          </NavLink>
        </li>
        <li className="left">
          <i className="bi bi-people icon"></i>
          <NavLink to="/users" className="nav-link">
            Users
          </NavLink>
        </li>
      </ul>
    </nav>
  );
}
export default Sidebar;
