import { Link } from "react-router-dom";
import { useAuthStore } from "../store/useAuthStore";
import { LogOut, MessageSquare, Settings, User } from "lucide-react";

const Navbar = () => {
  const { logout, authUser } = useAuthStore();

  return (
    <header className="bg-base-100/80 border-b border-base-300 fixed w-full backdrop-blur-lg">
      <div className="container flex flex-row justify-between items-center gap-2 pt-2 pb-3">
        {/*Home Link*/}
        <Link to="/" className="text-center group ml-5">
          <div className="flex flex-row justify-center items-center gap-2">
            <div
              className="flex justify-center items-center size-8 bg-primary/10 
            group-hover:bg-primary/40 group-active:bg-primary/70 rounded-xl transition-colors"
            >
              <MessageSquare className="size-4 text-primary" />
            </div>
            <span className="font-bold text-lg group-hover:text-xl group-active:text-lg">
              Chatty
            </span>
          </div>
        </Link>

        <div className="flex flex-row gap-3 mr-4">
          {/*Settings Link*/}
          <Link to="/settings">
            <button className="btn btn-soft">
              <Settings className="size-3" />
              <span className="text-xs">Settings</span>
            </button>
          </Link>

          {/*Profile Link*/}
          {authUser && (
            <Link to="/profile">
              <button className="btn btn-soft">
                <User className="size-3" />
                <span className="text-xs">Profile</span>
              </button>
            </Link>
          )}

          {/*Logout Link*/}
          {authUser && (
            <button className="btn btn-soft" onClick={logout}>
              <LogOut className="size-3" />
              <span className="text-xs">Logout</span>
            </button>
          )}
        </div>
      </div>
    </header>
  );
};

export default Navbar;
