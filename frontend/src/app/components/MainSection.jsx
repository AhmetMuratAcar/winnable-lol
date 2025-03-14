"use client"
import Image from "next/image";

const MainSection = () => {
	return (
		<section id="MainSection" className="flex flex-col flex-grow items-center justify-start">
			<Image 
				src="/images/cope.png"
				width={718}
				height={347}
				alt="copium pepe"
				className="m-10 rounded-4xl"
			/>

			<form className="flex items-center bg-[#31313c] rounded-full px-4 py-2 w-full max-w-xl">
				{/* <!-- Region Dropdown --> */}
				<div className="ml-4 flex flex-col">
					<label className="text-xs text-white font-semibold mb-1 ml-1">Region</label>
					<select className="bg-transparent text-gray-400 outline-none">
						<option>North America</option>
						<option>Europe West</option>
						<option>Europe Nordic & East</option>
						<option>Middle East</option>
						<option>Oceania</option>
						<option>LAS</option>
						<option>LAN</option>
						<option>Southeast Asia</option>
						<option>Korea</option>
						<option>Japan</option>
						<option>Brazil</option>
						<option>Russia</option>
						<option>Türkiye</option>
						<option>Taiwan</option>
						<option>Vietnam</option>
					</select>
				</div>

				{/* <!-- Divider --> */}
				<span className="w-px h-6 bg-gray-500 mx-4"></span>

				{/* <!-- Search Input --> */}
				<div className="flex flex-col flex-grow">
					<label className="text-xs text-white font-semibold mb-1">Search</label>
					<input
						type="text"
						placeholder="IGN + #Tag"
						className="bg-transparent text-white outline-none placeholder-gray-400"
					/>
				</div>
			</form>
			
			{/* <!-- Submit Button --> */}
			<button type="submit" className="mt-5 px-4 py-2 rounded-md bg-[var(--contrast)] text-white font-bold hover:opacity-80 transition">
				Winnable?
			</button>
		</section>
	);
};

export default MainSection;
