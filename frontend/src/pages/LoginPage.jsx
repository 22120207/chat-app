import { useState } from "react";
import toast from "react-hot-toast";
import { Link } from "react-router-dom";
import {
  Eye,
  EyeOff,
  Loader2,
  KeyRound,
  MessageSquare,
  User,
} from "lucide-react";

import { useAuthStore } from "../store/useAuthStore";
import AuthImagePattern from "../components/AuthImagePattern";

const LoginPage = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });

  const { login, isLoggingIn } = useAuthStore();

  const validateForm = () => {
    if (!formData.username.trim()) return toast.error("Username is required");

    if (!formData.password) return toast.error("Password is required");

    if (formData.password.length < 6)
      return toast.error("Password must be at least 6 characters");

    return true;
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    const success = validateForm();

    if (success === true) login(formData);
  };

  return (
    <div className="min-h-screen grid lg:grid-cols-2 sm:grid-cols-1 select-none">
      {/* Left side */}
      <div className="flex flex-col justify-center items-center">
        <div className="w-full max-w-md space-y-8 bg-base-200 rounded-2xl py-5 px-7">
          {/* Logo */}
          <div className="text-center mb-8">
            <div className="flex flex-col items-center gap-2">
              <div
                className="flex justify-center items-center size-12 bg-primary/10 
            hover:bg-primary/40 active:bg-primary/70 rounded-xl transition-colors"
              >
                <MessageSquare className="size-6 text-primary" />
              </div>
              <h1 className="text-2xl font-bold">Welcome Back!</h1>
              <p className="text-base-content/60 text-sm">
                Sign in to your account
              </p>
            </div>
          </div>

          {/* Form */}
          <form onSubmit={handleSubmit}>
            <div className="flex flex-col gap-8">
              {/* Username Input */}
              <label className="input w-full">
                <User className="size-5 text-base-content/60" />
                <input
                  type="text"
                  placeholder="Username"
                  value={formData.username}
                  onChange={(e) =>
                    setFormData({ ...formData, username: e.target.value })
                  }
                />
              </label>

              {/* Password Input */}
              <label className="input w-full">
                <KeyRound className="size-5 text-base-content/60" />
                <input
                  type={showPassword ? "" : "password"}
                  placeholder="Password"
                  value={formData.password}
                  onChange={(e) =>
                    setFormData({ ...formData, password: e.target.value })
                  }
                />
                <button
                  type="button"
                  className="hover:text-base-content"
                  onClick={() => {
                    setShowPassword(!showPassword);
                  }}
                >
                  {showPassword ? (
                    <Eye className="size-5 text-base-content/60" />
                  ) : (
                    <EyeOff className="size-5 text-base-content/60" />
                  )}
                </button>
              </label>
            </div>

            {/* Create Button */}
            <button
              type="submit"
              className="btn btn-primary w-full mt-8 mb-2"
              disabled={isLoggingIn}
            >
              {isLoggingIn ? (
                <>
                  {" "}
                  <Loader2 className="animate-spin" />
                  Loading...
                </>
              ) : (
                "Sign In"
              )}
            </button>

            {/* Link to Login */}
            <p className="text-center text-base-content/60 mt-0.5">
              Don't have an account?{" "}
              <Link className="link link-primary" to="/signup">
                Sign Up
              </Link>
            </p>
          </form>
        </div>
      </div>

      {/* Right side */}
      <AuthImagePattern
        title={"Welcome Back!"}
        subtitle={
          "Sign in to continue your conversations and catch up with your messages."
        }
      />
    </div>
  );
};

export default LoginPage;
