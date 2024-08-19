let currentParagraph = "bio";

let prevParagraph;
let nextParagraph;

function changeParagraph(newParagraph) {
    if (currentParagraph != newParagraph) {
        prevParagraph = document.getElementById(currentParagraph);
        nextParagraph = document.getElementById(newParagraph);

        currentParagraph = newParagraph;
        prevParagraph.classList.add("fade-out");

        prevParagraph.addEventListener('animationend', () => {
            prevParagraph.classList.add("hidden");
            prevParagraph.classList.remove("fade-out");

            nextParagraph.classList.remove("hidden");
            nextParagraph.classList.add("fade-in");

            nextParagraph.addEventListener('animationend', () => {
            nextParagraph.classList.remove("fade-in");
            })
        })
    }
}