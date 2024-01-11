import { Outlet } from "react-router-dom";
import Sidebar from "./Sidebar";
import "./Admin.css";

function Admin() {
  return (
    <>
      <div className="wrapper">
        <Sidebar></Sidebar>
        <div id="border"></div>
        <div className="content">
          <Outlet />
        </div>
      </div>
    </>
  );
}
export default Admin;
