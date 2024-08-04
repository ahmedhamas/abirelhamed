const id = window.location.pathname.split("/")[2];

async function getCase() {
  const response = await fetch(`/api/case/${id}`);
  const data = await response.json();

  console.log(data);

  document.querySelector("#case").innerHTML = `
    <ul class="list-group">

    <li class="list-group-item"><strong>معرف:</strong> ${data.Id}</li>
    <li class="list-group-item"><strong>اسم الحالة:</strong> ${
      data.Case_name
    }</li>
    <li class="list-group-item"><strong>الرقم القومي:</strong> ${
      data.National_id.String
    }</li>
    <li class="list-group-item"><strong>أجهزة تحتاجها الحالة:</strong> ${
      data.Devices_needed_for_the_case
    }</li>
    <li class="list-group-item"><strong>اجمالي الدخل:</strong> ${
      data.Total_income
    }</li>
    <li class="list-group-item"><strong>مصروفات ثابتة:</strong> ${
      data.Fixed_expenses
    }</li>
    <li class="list-group-item"><strong>معاش عن الزوج:</strong> ${
      data.Pension_from_husband
    }</li>
    <li class="list-group-item"><strong>معاش عن الأب:</strong> ${
      data.Pension_from_father
    }</li>
    <li class="list-group-item"><strong>الديون:</strong> ${data.Debts}</li>
    <li class="list-group-item"><strong>نوع الحالة:</strong> ${
      data.Case_type
    }</li>
    <li class="list-group-item"><strong>تاريخ الميلاد:</strong> ${
      data.Date_of_birth.String.slice(0, 4) +
      "/" +
      data.Date_of_birth.String.slice(5, 7) +
      "/" +
      data.Date_of_birth.String.slice(8, 10)
    }</li>
    <li class="list-group-item"><strong>العمر:</strong> ${data.Age}</li>
    <li class="list-group-item"><strong>النوع:</strong> ${data.Gender}</li>
    <li class="list-group-item"><strong>الحالة الاجتماعية:</strong> ${
      data.Social_situation
    }</li>
    <li class="list-group-item"><strong>العنوان من بطاقة الهوية:</strong> ${
      data.Address_from_national_id_card
    }</li>
    <li class="list-group-item"><strong>العنوان الفعلي:</strong> ${
      data.Actual_address
    }</li>
    <li class="list-group-item"><strong>حي:</strong> ${data.District}</li>
    <li class="list-group-item"><strong>تاريخ دخول الحالة:</strong> ${
      data.Created_at.String.slice(0, 4) +
      "/" +
      data.Created_at.String.slice(5, 7) +
      "/" +
      data.Created_at.String.slice(8, 10)
    }</li>
    <li class="list-group-item"><strong>تاريخ تجديد البحث:</strong> ${
      data.Updated_at.String.slice(0, 4) +
      "/" +
      data.Updated_at.String.slice(5, 7) +
      "/" +
      data.Updated_at.String.slice(8, 10)
    }</li>

    </ul>
     <h4 style="margin-top: 20px">بيانات الزوج</h4>
     <div class="sonsbuttons">
         <button class="edit-button" onclick="edit()">
                <i class='bx bxs-edit-alt'></i>
              </button>
              <button class="add-button" onclick="add()">
                <i class='bx bx-plus'></i>
              </button>
    </div>
  <ul class="list-group">
    <li class="list-group-item"><strong>اسم الزوج:</strong> ${
      data.Husband_name
    }</li>
    <li class="list-group-item">
      <strong>الرقم القومي:</strong> ${data.Husband_national_id.String}
    </li>
    <li class="list-group-item">
      <strong>تاريخ الميلاد:</strong> ${
        data.Husband_date_of_birth.String.slice(0, 4) +
        "/" +
        data.Husband_date_of_birth.String.slice(5, 7) +
        "/" +
        data.Husband_date_of_birth.String.slice(8, 10)
      }
    </li>
    <li class="list-group-item"><strong>العمر:</strong> ${data.Husband_age}</li>
    <li class="list-group-item"><strong>النوع:</strong> ${
      data.Husband_gender
    }</li>
  </ul>
   <h4 style="margin-top: 20px">الوضع الاجتماعي</h4>
   <div class="sonsbuttons">
         <button class="edit-button" onclick="editCase()">
                <i class='bx bxs-edit-alt'></i>
              </button>
               <button class="add-button" onclick="add()">
                <i class='bx bx-plus'></i>
              </button>
    </div>
  <ul class="list-group">
    <li class="list-group-item"><strong>خصائص:</strong> ${data.Properties}</li>
    <li class="list-group-item">
      <strong>الحالة الصحية:</strong> ${data.Health_status}
    </li>
    <li class="list-group-item"><strong>التعليم:</strong> ${data.Education}</li>
    <li class="list-group-item">
      <strong>عدد أفراد الأسرة:</strong> ${data.Number_of_family_members}
    </li>
    <li class="list-group-item">
      <strong>عدد الأطفال المسجلين:</strong> ${
        data.Number_of_registered_children
      }
    </li>
    <li class="list-group-item">
      <strong>إجمالي عدد الأطفال:</strong> ${data.Total_number_of_children}
    </li>
  </ul>
  <h4 style="margin-top: 20px">الإعانات</h4>
  <div class="sonsbuttons">
         <button class="edit-button" onclick="editCase()">
                <i class='bx bxs-edit-alt'></i>
              </button>
               <button class="add-button" onclick="add()">
                <i class='bx bx-plus'></i>
              </button>
    </div>
  <ul class="list-group">
    <li class="list-group-item">
      <strong>منح من خارج الجمعية:</strong>
      ${data.Grants_from_outside_the_association}
    </li>
    <li class="list-group-item">
      <strong>منح مالية من خارج الجمعية:</strong>
      ${data.Grants_from_outside_the_association_financial}
    </li>
    <li class="list-group-item">
      <strong>منح مالية من الجمعية:</strong>
      ${data.Grants_from_the_association_financial}
    </li>
    <li class="list-group-item">
      <strong>منح عينية من الجمعية:</strong>
      ${data.Grants_from_the_association_financial}
    </li>
    <li class="list-group-item">
      <strong>إجمالي الإعانات:</strong> ${data.Total_Subsidies}
    </li>
  </ul><h4 style="margin-top: 20px">الأبناء والأقارب</h4>
    <div class="sons" style="
    display: flex;
    flex-direction: column;
    gap: 20px;
">
${data.Relatives.map((rel) => {
  return ` 
      <div class="sonsbuttons">
         <button class="edit-button" onclick="editCase()">
                <i class='bx bxs-edit-alt'></i>
              </button>
               <button class="add-button" onclick="add()">
                <i class='bx bx-plus'></i>
              </button>
    </div>
  <ul class="list-group">
    <li class="list-group-item">
      <strong>نوع القريب:</strong> ${rel.Relative_type.String}
    </li>
    <li class="list-group-item"><strong>اسم القريب:</strong> ${
      rel.Relative_name
    }</li>
    <li class="list-group-item">
      <strong>الرقم القومي:</strong> ${rel.Relative_national_id.String}
    </li>
    <li class="list-group-item">
      <strong>تاريخ الميلاد:</strong> ${
        rel.Relative_date_of_birth.String.slice(0, 4) +
        "/" +
        rel.Relative_date_of_birth.String.slice(5, 7) +
        "/" +
        rel.Relative_date_of_birth.String.slice(8, 10)
      }
    </li>
    <li class="list-group-item"><strong>العمر:</strong> ${rel.Relative_age}</li>
    <li class="list-group-item"><strong>النوع:</strong> ${
      rel.Relative_gender
    }</li>
    <li class="list-group-item"><strong>الوظيفة:</strong> ${
      rel.Relative_job
    }</li>
    <li class="list-group-item">
      <strong>الوضع الاجتماعي:</strong> ${rel.Relative_social_situation}
    </li>
    <li class="list-group-item">
      <strong>الحالة الصحية:</strong> ${rel.Relative_health_status}
    </li>
    <li class="list-group-item"><strong>التعليم:</strong> ${
      rel.Relative_education
    }</li>
  </ul>
  `;
}).join("")}
  </div>

  `;
}

getCase();
