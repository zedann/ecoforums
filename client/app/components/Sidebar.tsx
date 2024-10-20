const Sidebar = () => {
  return (
    <div className="bg-gray-100 p-4 rounded-lg shadow-md my-4">
      <h2 className="text-2xl font-bold mb-4 text-black ">Trendings</h2>
      <ul>
        <li className="mb-2">
          <a href="#" className="text-blue-500">
            Link 1
          </a>
        </li>
        <li className="mb-2">
          <a href="#" className="text-blue-500">
            Link 2
          </a>
        </li>
        <li className="mb-2">
          <a href="#" className="text-blue-500">
            Link 3
          </a>
        </li>
      </ul>
    </div>
  );
};

export default Sidebar;
