import { Outlet } from "react-router-dom";
import Sidebar from "./Sidebar";
import "./Admin.css";

function Admin() {
  return (
    <>
      <div className="container">
        <Sidebar></Sidebar>
        <div className="content">
          <Outlet />
        </div>
      </div>
    </>
  );
}
export default Admin;
