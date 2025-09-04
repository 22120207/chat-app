import { MessageSquare } from "lucide-react";

const NoChatSelected = () => {
  return (
    <div className="w-full flex flex-1 items-center justify-center select-none">
      <div className="text-center mb-5 lg:mb-8">
        <div className="flex flex-col items-center gap-1 lg:gap-2">
          <div
            className="flex justify-center items-center size-12 bg-primary/10 
            hover:bg-primary/40 active:bg-primary/70 rounded-xl transition-colors animate-bounce"
          >
            <MessageSquare className="size-6 text-primary" />
          </div>
          <h1 className="text-lg md:text-2xl font-bold">Welcome to Chatty!</h1>
          <p className="text-base-content/60 text-xs md:text-sm">
            Select a conversation from the sidebar to start chatting
          </p>
        </div>
      </div>
    </div>
  );
};

export default NoChatSelected;
