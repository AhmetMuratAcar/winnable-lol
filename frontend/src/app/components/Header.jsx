"use client";
import React, { useState, useRef, useEffect } from 'react';
import { FaQuestionCircle } from 'react-icons/fa';
import Image from 'next/image';

const Header = () => {
  const [showInfo, setShowInfo] = useState(false);
  const infoRef = useRef(null);

  // Close the info box when clicking outside of it
  useEffect(() => {
    function handleClickOutside(event) {
      // If the popup is open and the click is not inside infoRef, close it
      if (showInfo && infoRef.current && !infoRef.current.contains(event.target)) {
        setShowInfo(false);
      }
    }

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [showInfo]);

  return (
    <header className="w-full flex items-center justify-between px-4 bg-(--contrast) h-13">
      <div>
        <button 
          onClick={() => location.reload()}
          className="p-0 border-none bg-transparent cursor-pointer"
        >
          <Image
            src="/images/lotad-icon.png"
            width={45}
            height={45}
            alt="lotad"
            className="rounded-full"
          />
        </button>
      </div>
      <div className="relative" ref={infoRef}>
        <FaQuestionCircle
          className="text-2xl cursor-pointer"
          onClick={() => setShowInfo(!showInfo)}
        />
        {showInfo && (
          <div className="absolute right-0 mt-2 w-64 p-4 bg-white text-gray-800 border border-gray-300 rounded shadow-lg">
            <p className="text-sm">
              Welcome to Winnable. Enter your Riot ID to check the win probability prediction of your current game.
            </p>
          </div>
        )}
      </div>
    </header>
  );
};

export default Header;
