// const auth = localStorage.getItem("auth");

// if (!auth) {
//     window.location.href = "/login";
//     localStorage.clear();
// }

async function getCases(limit = 30) {
  if (window.location.pathname !== "/") return;
  try {
    const response = await fetch(`/api?limit=${limit || 30}`);
    const data = await response.json();

    const pages = Math.ceil(data.Pages / 30);

    console.log(limit);

    const html = data.Cases.map((item) => {
      Object.keys(item).forEach((key) => {
        if (key === "date_of_birth") {
          if (item[key].String === "") return;
          const date = new Date(item[key].String);
          const year = date.getFullYear();
          const month = String(date.getMonth() + 1).padStart(2, "0");
          const day = String(date.getDate()).padStart(2, "0");
          item[key] = `${year}/${month}/${day}`;
        }
      });

      console.log(item);

      return `
        <tr>
            <td>${item.id}</td>
            <td>${item.case_name}</td>
            <td>${item.national_id}</td>
            <td>${item.devices_needed_for_the_case}</td>
            <td>${item.total_income}</td>
            <td>${item.fixed_expenses}</td>
            <td>${item.pension_from_husband}</td>
            <td>${item.pension_from_father}</td>
            <td>${item.debts}</td>
            <td>${item.case_type}</td>
            <td>${item.date_of_birth}</td>
            <td>${item.age}</td>
            <td>${item.gender}</td>
            <td>${item.social_situation}</td>
            <td>${item.address_from_national_id_card}</td>
            <td>${item.actual_address}</td>
            <td>${item.district}</td>
            <td>${
              item.created_at.String.slice(0, 4) +
              "/" +
              item.created_at.String.slice(5, 7) +
              "/" +
              item.created_at.String.slice(8, 10)
            }</td>
            <td>${
              item.updated_at.String.slice(0, 4) +
              "/" +
              item.updated_at.String.slice(5, 7) +
              "/" +
              item.updated_at.String.slice(8, 10)
            }</td>
            <td>
              <button class="edit-button" onclick="editCase(${Object.keys(item)
                .map((key) => {
                  if (key === "created_at" || key === "updated_at") {
                    return item[key].String;
                  }
                  return item[key];
                })
                .join("|")})">
                <i class='bx bxs-edit-alt'></i>
              </button>
              <button class="delete-button" onclick="deleteCase(${item.id})">
                <i class='bx bxs-trash'></i>
              </button>
              <button class="view-button" onclick="GoTo(${item.id})">
                <i class='bx bxs-show'></i>
              </button>
            </td>
        </tr>
    `;
    }).join("");

    const table = document.querySelector(".casesBody");
    table.innerHTML = html;
    const page = Math.ceil(limit / 30);
    createPagination(pages, page);
  } catch (error) {
    console.error(error);
  }
}

getCases();
function createPagination(totalPages, currentPage) {
  const pagination = document.getElementById("pagination");
  pagination.innerHTML = "";

  const maxPageButtons = 5;

  const createPageItem = (page, text = page) => {
    const li = document.createElement("li");
    li.className = `page-item ${page === currentPage ? "active" : ""}`;
    const button = document.createElement("button");
    button.className = "page-link";
    button.textContent = text;
    button.addEventListener("click", () => updatePagination(page));
    li.appendChild(button);
    return li;
  };

  const updatePagination = (newPage) => {
    currentPage = newPage;
    getCases(newPage * 30);
    createPagination(totalPages, currentPage);
  };

  if (currentPage > 1) {
    pagination.appendChild(createPageItem(currentPage - 1, "Previous"));
  }

  let startPage = Math.max(1, currentPage - Math.floor(maxPageButtons / 2));
  let endPage = Math.min(totalPages, startPage + maxPageButtons - 1);

  if (startPage > 1) {
    pagination.appendChild(createPageItem(1));
    const ellipsis = document.createElement("li");
    ellipsis.className = "page-item disabled";
    ellipsis.innerHTML = '<button class="page-link">...</button>';
    pagination.appendChild(ellipsis);
  }

  for (let i = startPage; i <= endPage; i++) {
    pagination.appendChild(createPageItem(i));
  }

  if (endPage < totalPages) {
    const ellipsis = document.createElement("li");
    ellipsis.className = "page-item disabled";
    ellipsis.innerHTML = '<button class="page-link">...</button>';
    pagination.appendChild(ellipsis);
    pagination.appendChild(createPageItem(totalPages));
  }

  if (currentPage < totalPages) {
    pagination.appendChild(createPageItem(currentPage + 1, "Next"));
  }
}

function GoTo(id) {
  window.location.href = `/case/${id}`;
}

function AddForm(open) {
  if (open) {
    document.getElementById("add").style.display = "flex";
  } else {
    document.getElementById("add").style.display = "none";
  }
}

document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("add");

  form.addEventListener("submit", async (event) => {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);

    const data = Object.fromEntries(formData);

    const response = await fetch("/cases/add", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (response.ok) {
      getCases();
    }
  });
});

function deleteCase(id) {
  fetch(`/case/${id}`, {
    method: "DELETE",
  }).then((response) => {
    if (response.ok) {
      getCases();
    }
  });
}
