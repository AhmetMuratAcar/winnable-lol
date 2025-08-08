import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";
import ProfileSection from '@/app/components/ProfileSection';

export default async function SummonerPage ({ params }) {
    const { region, slug } = await params

    return (
        <main className="min-h-svh flex flex-col">
            <Header />
            <ProfileSection region={region} slug={slug}/>
            <Footer />
        </main>
    );
}