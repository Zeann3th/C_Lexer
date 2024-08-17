#include <bits/stdc++.h>
#include <iostream>
#include <limits>
using namespace std;
// Những biến toàn cục
int dem = 0;
vector<string> a;
// Function
void addToList() {
  cout << "Describe your activity" << endl;
  dem++;
  string s;
  cin.ignore(numeric_limits<streamsize>::max(), '\n');
  getline(cin, s);
  cin.clear();
  a.push_back(s);
}
void deleteFromList() {
  int choice;
  cout << "Select an activity to delete from the list" << endl;
  cout << "(If you wish to exit, please enter 0)" << endl;
  cin >> choice;
  while (choice > dem) {
    cout << "Please enter an appropriate number: "
         << endl; // Khi "cố tình" chọn số lớn hơn số có trong list
    cin >> choice;
  }
  a.erase(a.begin() + choice - 1); // xóa phần tử đã chọn
  dem--;
}
void Save() {
  ofstream myLog;
  myLog.open("Log.txt", ios::out); // Có thể sửa tên file ở đây
  if (myLog.fail()) {
    cout << "No Log file" << endl;
  } else {
    for (int i = 0; i < a.size(); i++) {
      myLog << a[i] << endl;
    }
    myLog.close();
  }
}
void loadDataFile() {
  ifstream myLog;
  myLog.open("Log.txt", ios::in); // Có thể sửa tên file ở đây
  if (myLog.fail()) {
    cout << "No Log file" << endl;
  } else {
    dem = 0; // reset lại đếm để in bảng mới
    string b;
    while (!myLog.eof()) {
      getline(myLog, b); // Đọc từng dòng trong file
      if (myLog.good()) {
        a.push_back(b); // đẩy nội dung các dòng trong file vào các phần tử
                        // tương ứng trong vector
        dem++; // đếm số phần tử để in trong main
      }
    }
    myLog.close();
  }
}
// Main ??
int main() {
  int flag = 1; // Quyết định file tiếp tục hay đóng
  int choice;
  do {
    system("cls"); // Reset lại màn hình
    cout << "--------------------To-do List--------------------" << endl;
    if (dem == 0) {
      cout << "----------------------(EMPTY)---------------------" << endl;
      cout << endl;
    } else {
      for (int i = 0; i < dem; i++) {
        cout << i + 1 << ". " << a[i] << endl;
      }
      cout << "--------------------------------------------------" << endl;
    }
    cout << "1. Add an activity" << endl;
    cout << "2. Delete an activity" << endl;
    cout << "3. Load data from existing file" << endl;
    cout << "4. Save" << endl;
    cout << "5. Exit" << endl;
    cin >> choice;
    switch (choice) {
    case 1:
      addToList();
      break;
    case 2:
      deleteFromList();
      break;
    case 3:
      loadDataFile();
      break;
    case 4:
      Save();
      break;
    case 5:
      flag = 0; // Exit();
      break;
    default: // Vẫn đang lỗi khi nhập chữ, số thập phân. Nhập từ 0-9 thì không
             // bị lỗi
      cout << "Please enter a number from 1-5" << endl;
      system("pause"); // solution hiện tại, vì nếu ko có thì hệ thống sẽ lại
                       // reset, ko hiện cảnh báo nữa
      break;
    }
  } while (flag == 1);
  return 0;
}
