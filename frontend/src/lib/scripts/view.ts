export enum Views {
  Verify = "verify",
  DashHome = "dashboard-home",
  DashColor = "dashboard-color-picker",
}

const possibleViews: string[] = [
  "verify",
  "dashboard-home",
  "dashboard-color-picker",
];

const viewStorageId: string = "current-view";

export const changeView = (view: Views) => {
  if (possibleViews.includes(view)) {
    if (window) {
      window.localStorage.setItem(viewStorageId, view);
    }
  }
};

export const getCurrentView = (): string => {
  if (window) {
    const view: string | null = window.localStorage.getItem(viewStorageId);
    if (view != null) {
      return view;
    }
  }
};
