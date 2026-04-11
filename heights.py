import re
with open("tmp_mystery/mystery/vehicles") as f:
    v = f.read()
clubs = [
    set([l.strip() for l in open("tmp_mystery/mystery/memberships/AAA").readlines() if l.strip()]),
    set([l.strip() for l in open("tmp_mystery/mystery/memberships/Delta_SkyMiles").readlines() if l.strip()]),
    set([l.strip() for l in open("tmp_mystery/mystery/memberships/Museum_of_Bash_History").readlines() if l.strip()]),
    set([l.strip() for l in open("tmp_mystery/mystery/memberships/Terminal_City_Library").readlines() if l.strip()])
]
in_all = clubs[0] & clubs[1] & clubs[2] & clubs[3]
print("All 4 Memberships:")
for owner in in_all:
    idx = v.find("Owner: " + owner)
    if idx > -1:
        block = v[idx:idx+150]
        make_color = "Unknown"
        h = re.search(r"Height: (.*)", block)
        print(owner, "-", h.group(1) if h else "?")
for owner in ["Nikolaus Milatz", "Augustin Lozano", "Matt Waite", "Brian Boyer", "Mike Bostock", "Dartey Henv"]:
    idx = v.find("Owner: " + owner)
    block = v[max(idx-100, 0):idx+150]
    p = re.search(r"License Plate (.*)", block)
    print(owner, p.group(1))
