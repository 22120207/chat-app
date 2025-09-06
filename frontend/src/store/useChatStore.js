import { create } from "zustand";

import { axiosInstance } from "../lib/axios";
import toast from "react-hot-toast";

export const useChatStore = create((set, get) => ({
  messages: [],
  users: [],
  selectedUser: null,
  isUsersLoading: false,
  isMessagesLoading: false,
  socket: null,

  connectSocket: () => {
    if (get().socket) return;

    const socket = new WebSocket("ws://localhost:5000/api/ws");

    socket.onopen = () => {
      console.log("WebSocket connected");
    };

    socket.onclose = () => {
      console.log("WebSocket disconnected");
      set({ socket: null });
    };

    set({ socket });
  },

  disconnectSocket: () => {
    const socket = get().socket;

    if (socket) {
      socket.close();
      set({ socket: null });
    }
  },

  getUsers: async () => {
    set({ isUsersLoading: true });
    try {
      const res = await axiosInstance.get("/users");
      set({ users: res.data });
    } catch (error) {
      toast.error(error.response.data.message);
    } finally {
      set({ isUsersLoading: false });
    }
  },

  getMessages: async (userId) => {
    set({ isMessagesLoading: true });
    try {
      const res = await axiosInstance.get(`/messages/${userId}`);
      set({ messages: res.data });
    } catch (error) {
      toast.error(error.response.data.message);
    } finally {
      set({ isMessagesLoading: false });
    }
  },

  sendMessage: async (messageData) => {
    const { selectedUser, messages, socket } = get();
    try {
      const res = await axiosInstance.post(
        `/messages/send/${selectedUser.id}`,
        messageData
      );

      set({ messages: [...messages, res.data] });

      if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(
          JSON.stringify({
            receiverId: selectedUser.id,
            message: messageData.message,
          })
        );
      }
    } catch (error) {
      toast.error(error.response?.data?.message || "Failed to send message");
    }
  },

  subscribeToMessages: () => {
    const socket = get().socket;
    if (!socket) return;

    socket.onmessage = (event) => {
      try {
        const newMessage = JSON.parse(event.data);

        const { selectedUser } = get();
        const isMessageSentFromSelectedUser =
          newMessage.senderId === selectedUser?._id;

        if (isMessageSentFromSelectedUser) {
          set({ messages: [...get().messages, newMessage] });
        }
      } catch (err) {
        console.error("Failed to parse WS message", err);
      }
    };
  },

  unsubscribeFromMessages: () => {
    const socket = get().socket;
    if (socket) {
      socket.onmessage = null;
    }
  },

  setSelectedUser: (selectedUser) => set({ selectedUser }),
}));
