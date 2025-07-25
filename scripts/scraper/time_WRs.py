import requests
import re
import json
import os


def fetch_data(url: str) -> dict[str, float]:
    headers = {
        "User-Agent": "Mozilla/5.0",
        "Accept-Language": "en-US,en;q=0.9"
    }
    res = requests.get(url, headers=headers)
    html = res.text

    match = re.search(r"ChampionsPage\.init\((\{.*?\})\);", html, re.DOTALL)
    if not match:
        raise ValueError("Could not find ChampionsPage.init(...) in HTML")

    json_text = match.group(1)
    data = json.loads(json_text)

    wr_data = {}
    for champ in data["rankings"]:
        id = champ["popularity"]["championId"]
        win_rate = champ["popularity"]["winRate"]
        wr_data[id] = win_rate

    return wr_data


def write_json(data: dict[str, float], path: str) -> None:
    with open(path, 'w', encoding='utf-8') as f:
        json.dump(data, f, indent=2)
    print(f"JSON dumped to: {path}")


def main():
    urls = [
        "https://www.leagueofgraphs.com/champions/tier-list/short",
        "https://www.leagueofgraphs.com/champions/tier-list/medium",
        "https://www.leagueofgraphs.com/champions/tier-list/long"
    ]

    write_dir = "/Users/murat/winnable-lol/data"

    for url in urls:
        data = fetch_data(url=url)
        file_name = url.split("/")[-1] + "_wr.json"
        total_path = os.path.join(write_dir, file_name)
        write_json(data=data, path=total_path)


if __name__ == "__main__":
    main()
