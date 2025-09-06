import { useChatStore } from "../store/useChatStore";

import ChatHeader from "./ChatHeader";
import MessageInput from "./MessageInput";
import MessageSkeleton from "./skeletons/MessageSkeleton";
import { useAuthStore } from "../store/useAuthStore";
import { formatMessageTime } from "../lib/utils";
import { useEffect } from "react";

const ChatContainer = () => {
  const { messages, getMessages, isMessagesLoading, selectedUser } =
    useChatStore();

  const { authUser } = useAuthStore();

  useEffect(() => {
    getMessages(selectedUser.id);
  }, [selectedUser.id, getMessages]);

  if (isMessagesLoading) {
    return (
      <div className="flex-1 flex flex-col overflow-auto">
        <ChatHeader />
        <MessageSkeleton />
        <MessageInput />
      </div>
    );
  }

  return (
    <div className="flex-1 flex flex-col overflow-auto">
      <ChatHeader />

      <div className="flex-1 overflow-y-auto p-4 space-y-4">
        {messages.length > 0
          ? messages.map((message, index) => (
              <div
                key={index}
                className={`chat ${
                  message.SenderID === authUser.id ? "chat-end" : "chat-start"
                }`}
              >
                <div className="chat-image avatar">
                  <div className="size-10 rounded-full border">
                    <img
                      src={
                        message.SenderID === authUser.id
                          ? authUser.profilePic
                          : selectedUser.profilePic
                      }
                    />
                  </div>
                </div>
                <div className="chat-header mb-1">
                  <time className="text-xs opacity-50 ml-1">
                    {formatMessageTime(message.CreatedAt)}
                  </time>
                </div>
                <div className="chat-bubble flex flex-col">
                  {message.Message && <p>{message.Message}</p>}
                </div>
              </div>
            ))
          : ""}
      </div>

      <MessageInput />
    </div>
  );
};

export default ChatContainer;
