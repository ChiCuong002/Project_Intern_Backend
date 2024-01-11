function Users() {
  return (
    <>
      <div className="title">
        <h2>Users</h2>
      </div>
      <div className="nav-search">
        <input placeholder="Search Users" />
        <button type="submit">Go</button>
        <span>
          Showing results<span>1</span> to <span>5</span>
        </span>
      </div>
    </>
  );
}
export default Users;
