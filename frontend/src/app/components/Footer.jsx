export default function Footer({ bgClass = "bg-(--contrast)" }) {
  return (
    <section id="FooterSection">
      <footer
        className={`${bgClass} h-12 items-center justify-center text-center flex`}
      >
        <div>
          <p>
            Check it out on{" "}
            <a
              href="https://github.com/AhmetMuratAcar/winnable-lol"
              target="_blank"
              className="text-(--pastel-red) hover:underline"
            >
              Github
            </a>
          </p>
        </div>
      </footer>
    </section>
  );
}
