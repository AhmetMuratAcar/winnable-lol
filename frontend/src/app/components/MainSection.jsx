"use client"
import { useState } from "react";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { idValidation } from "../utils/idValidation";

export default function MainSection() {
	const router = useRouter()
	const [isSubmitting, setIsSubmitting] = useState(false)
	const [errorMessage, setErrorMessage] = useState('')
	
	async function onSubmit(event) {
		event.preventDefault()
		setIsSubmitting(true)
		setErrorMessage('')

		if (!navigator.onLine) {
			setErrorMessage('You appear to be offline, Please check your internet connection')
			setIsSubmitting(false)
			return
		}

		const formData = new FormData(event.currentTarget)
		for (const [key, value] of formData.entries()) {
			console.log(`${key}: ${value}`)
		}

		const validatedID = idValidation({
			region: event.target.region.value,
			riotID: event.target.ign.value
		})

		if (!validatedID.isValid) {
			const msg = 'Invalid riotID'
			setErrorMessage(msg)
			setIsSubmitting(false)
			return
		}

		const slug = `${encodeURIComponent(validatedID.gameName)}-${encodeURIComponent(validatedID.tagLine)}`
		router.push(`lol/summoners/${validatedID.region}/${slug}`)
		setIsSubmitting(false)
		return
	}

	return (
		<section
			id="MainSection"
			className="flex flex-col flex-grow items-center justify-start px-4 py-6"
		>
			<div className="relative w-[80%] max-w-[350px] m-10 mt-0 aspect-square">
			  <Image
			    src="/images/logo.png"
			    alt="logo"
			    fill
			    style={{ objectFit: 'contain' }}
			    sizes="(max-width: 640px) 80vw, 350px"
			    priority
			  />
			</div>

			<form 
				onSubmit={onSubmit}
				className="flex flex-col items-center w-full max-w-xl"
			>
				<div className="flex items-center bg-[#31313c] rounded-full px-4 py-2 w-full">
				
					{/* Region Dropdown */}
					<div className="w-24 sm:w-auto ml-4 flex flex-col shrink-0">
						<label 
							htmlFor="region" 
							className="text-xs text-white font-semibold mb-1 ml-1"
						>
							Region
						</label>
						<select 
							name="region"
							id="region"
							className="bg-transparent text-gray-400 outline-none"
							autoComplete="off"
						>
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

					{/* Divider */}
					<span className="w-px h-6 bg-gray-500 mx-4"></span>

					{/* Search Input */}
					<div className="flex flex-col flex-grow min-w-0">
						<label 
							htmlFor="ign" 
							className="text-xs text-white font-semibold mb-1"
						>
							Search
						</label>
						<input
							type="text"
							name="ign"
							id="ign"
							placeholder="IGN + #Tag"
							className="bg-transparent text-white outline-none placeholder-gray-400"
							required
						/>
					</div>
				</div>

				<button
					type="submit"
					className="mt-4 px-6 py-2 rounded-md bg-[var(--pastel-red)] text-white font-bold hover:opacity-80 transition"
					disabled={isSubmitting}
				>
					{isSubmitting ? 'Submitting…' : 'Winnable?'}
				</button>

				{errorMessage && (
					<p className="text-[var(--pastel-red)] text-sm mt-2">{errorMessage}</p>
				)}
			</form>
		</section>
  );
};