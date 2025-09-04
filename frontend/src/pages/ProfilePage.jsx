import { User, VenusAndMars } from "lucide-react";

import { useAuthStore } from "../store/useAuthStore";

const ProfilePage = () => {
  const { authUser } = useAuthStore();

  let gender = "male";

  if (authUser) {
    gender = authUser.gender;
  }

  gender = gender.charAt(0).toUpperCase() + gender.slice(1);

  return (
    <div className="min-h-screen flex items-center justify-center">
      <div
        className="flex flex-col items-center gap-3 justify-center w-full 
      max-w-md bg-base-200 rounded-2xl py-5 px-7"
      >
        <h1 className="text-lg md:text-2xl font-bold">Profile</h1>
        <p className="text-base-content/60 text-xs md:text-sm mb-2">
          <span className="font-bold">{authUser.fullname}</span> profile
          information
        </p>

        {/* Avatar */}
        <div className="avatar mb-4">
          <div className="w-24 rounded-full">
            <img src={authUser.profilePic} />
          </div>
        </div>

        {/* Line Seperate */}
        <div className="w-full border-1"></div>

        {/* Username display */}
        <label className="fieldset w-full mt-3">
          <legend className="flex flex-row gap-2">
            {<User className="size-5 text-base-content/60" />}
            {""}Username
          </legend>
          <label className="input mt-1 w-full">
            <p>{authUser.username}</p>
          </label>
        </label>

        {/* Gender display */}
        <label className="fieldset w-full">
          <legend className="flex flex-row gap-2">
            {<VenusAndMars className="size-5 text-base-content/60" />}
            {""}Gender
          </legend>
          <label className="input mt-1 w-full">
            <p>{gender}</p>
          </label>
        </label>

        {/* Account Information */}
        <div className="w-full m-2">
          <h4 className="text-sm font-bold mb-2">Account Information</h4>
          <div className="flex flex-row justify-between text-xs">
            <span>Member Since</span>
            <span className="italic">{authUser.createdAt?.split("T")[0]}</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProfilePage;
