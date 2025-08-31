export default function Footer() {
	return (
		<section id="FooterSection">
			<footer className="bg-(--contrast) h-12 items-center justify-center text-center flex">
				<div>
					<p>
						Check it out on{' '}
						<a 
							href="https://github.com/AhmetMuratAcar/winnable-lol"
							target='_blank'
							className="text-(--pastel-red) hover:underline"
						>Github
						</a>
					</p>
				</div>
			</footer>
		</section>
	);
};
