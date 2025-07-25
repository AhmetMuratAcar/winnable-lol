import sys
import json
from pathlib import Path


def validate_path(path_str: str) -> bool:
    path = Path(path_str)
    return (
        path.is_absolute()
        and path.suffix.lower() == '.json'
        and path.is_file()
        and path.exists()
    )


def map_tags(abs_path: str) -> dict[str, list[str]]:
    """
    Given the ABSOLUTE path to a champion.json or championFull.json file from
    a Riot dragontail release, maps the champion keys to a list of their tags.
    """
    try:
        with open(abs_path, 'r', encoding='utf-8') as f:
            data = json.load(f)
    except PermissionError:
        print(f"PermissionError: Cannot read file at {abs_path}")
        print("Check your file permissions")
        sys.exit(1)
    except json.JSONDecodeError:
        print(f"Invalid JSON: {abs_path}")
        sys.exit(1)

    res = {}
    champs = data['data']
    for champ in champs.values():
        res[champ['key']] = champ['tags']

    return res


def main():
    args = sys.argv[1:]

    if not args:
        print("ERROR: No file path provided")
        print('USAGE: python3 tag_mapper.py "<absolute path to champions.json>"')
        sys.exit(1)

    if len(args) > 1:
        print("ERROR: Too many inputs")
        print('USAGE: python3 tag_mapper.py "<absolute path to champions.json>"')
        sys.exit(1)

    path = args[0]
    if not validate_path(path):
        print("Error: Invalid path")
        print("Example: /Users/name/Downloads/dragontail-15.14.1/15.14.1/data/en_US/champion.json")

    tag_map = map_tags(abs_path=path)

    with open('tag_map.json', 'w', encoding='utf-8') as f:
        json.dump(tag_map, f, indent=2)


if __name__ == "__main__":
    main()
