import { useState } from "react";
import { useAuthStore } from "../store/useAuthStore";
import {
  Eye,
  EyeOff,
  Loader2,
  KeyRound,
  MessageSquare,
  User,
  VenusAndMars,
} from "lucide-react";
import { Link } from "react-router-dom";

import AuthImagePattern from "../components/AuthImagePattern";
import toast from "react-hot-toast";

const SignUpPage = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [formData, setFormData] = useState({
    fullName: "",
    username: "",
    password: "",
    confirmPassword: "",
    gender: "male",
  });

  const { signup, isSigningUp } = useAuthStore();

  const validateForm = () => {
    if (!formData.fullName.trim()) return toast.error("Full name is required");

    if (!formData.username.trim()) return toast.error("Username is required");

    if (!formData.password) return toast.error("Password is required");
    if (formData.password.length < 6)
      return toast.error("Password must be at least 6 characters");

    if (formData.password !== formData.confirmPassword) {
      return toast.error("Confirm Password must match Password exactly");
    }

    return true;
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    const success = validateForm();

    if (success === true) signup(formData);
  };

  return (
    <div className="min-h-screen grid lg:grid-cols-2 sm:grid-cols-1 select-none">
      {/* Left side */}
      <div className="flex flex-col justify-center items-center">
        <div className="w-full max-w-md space-y-8">
          {/* Logo */}
          <div className="text-center mb-8">
            <div className="flex flex-col items-center gap-2">
              <div
                className="flex justify-center items-center size-12 bg-primary/10 
            hover:bg-primary/40 active:bg-primary/70 rounded-xl transition-colors"
              >
                <MessageSquare className="size-6 text-primary" />
              </div>
              <h1 className="text-2xl font-bold">Create Account</h1>
              <p className="text-base-content/60">
                Get started with your free account
              </p>
            </div>
          </div>

          {/* Form */}
          <form onSubmit={handleSubmit}>
            <div className="flex flex-col gap-8">
              {/* Fullname Input */}
              <label class="input w-full">
                <input
                  id="fullname"
                  type="text"
                  placeholder="Your Full Name"
                  value={formData.fullName}
                  onChange={(e) =>
                    setFormData({ ...formData, fullName: e.target.value })
                  }
                />
              </label>

              {/* Username Input */}
              <label class="input w-full">
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
              <label class="input w-full">
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

              {/* Confirm password Input */}
              <label class="input w-full">
                <KeyRound className="size-5 text-base-content/60" />
                <input
                  type={showPassword ? "" : "password"}
                  placeholder="Confirm Password"
                  value={formData.confirmPassword}
                  onChange={(e) =>
                    setFormData({
                      ...formData,
                      confirmPassword: e.target.value,
                    })
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

              {/* Gender Select */}
              <select
                defaultValue="Select gender"
                className="select w-full"
                onChange={(e) =>
                  setFormData({ ...formData, gender: e.target.value })
                }
              >
                <option disabled={true}>Select gender</option>
                <option value="male">Male</option>
                <option value="female">Female</option>
              </select>
            </div>

            {/* Create Button */}
            <button
              type="submit"
              className="btn btn-primary w-full mt-8 mb-2"
              disabled={isSigningUp}
            >
              {isSigningUp ? (
                <>
                  {" "}
                  <Loader2 />
                  Loading...
                </>
              ) : (
                "Creating Account"
              )}
            </button>

            {/* Link to Login */}
            <p className="text-center text-base-content/60">
              Already have an account?{" "}
              <Link className="link link-primary" to="/login">
                Sign In
              </Link>
            </p>
          </form>
        </div>
      </div>

      {/* Right side */}
      <AuthImagePattern
        title={"Join our community"}
        subtitle="Connect with friends, share moments, and stay in touch with your loved ones."
      />
    </div>
  );
};

export default SignUpPage;
